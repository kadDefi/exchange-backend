package repo

import (
	"context"
	"exchange-backend/internal/domain"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func (r *Repo) UpdateMarket(ctx context.Context, ms *domain.Market) error {
	if err := r.db.WithContext(ctx).Model(ms).
		Where("id = ?", ms.ID).
		Updates(ms).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) CreateMarket(ctx context.Context, m *domain.Market) error {
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetMarketByAddress(ctx context.Context, collectionAddress string, tokenID string, status uint8) (*domain.Market, error) {
	market := new(domain.Market)
	if err := r.db.WithContext(ctx).
		Where("collection_address = ?", collectionAddress).
		Where("token_id = ?", tokenID).
		Where("status = ?", status).
		First(market).Error; err != nil {
		return nil, err
	}

	return market, nil
}

func (r *Repo) QueryMarket(ctx context.Context, arg *domain.QueryMarketArg) ([]*domain.Market, int, error) {
	var (
		list  []*domain.Market
		total int
	)

	tx := r.db.WithContext(ctx).Model(&domain.Market{})

	if len(arg.TokenIdIn) > 0 {
		tx = tx.Where("token_id = ?", arg.TokenIdIn)
	}

	if arg.Status != nil {
		tx = tx.Where("status = ?", *arg.Status)
	}

	if arg.SellerAddress != nil {
		tx = tx.Where("seller_address < ?", *arg.SellerAddress)
	}

	if arg.PriceGt != nil {
		tx = tx.Where("price_float > ?", *arg.PriceGt)
	}

	if arg.PriceLt != nil {
		tx = tx.Where("price_float < ?", *arg.PriceLt)
	}

	eg := errgroup.Group{}

	eg.Go(func() error {
		tx := tx.Session(&gorm.Session{})
		var count int64
		if err := tx.Count(&count).Error; err != nil {
			return err
		}

		total = int(count)

		return nil
	})

	eg.Go(func() error {
		tx := tx.Session(&gorm.Session{})
		if arg.Pagination.PerPage != nil {
			tx = tx.Limit(*arg.Pagination.PerPage)
			if arg.Pagination.Page != nil {
				tx = tx.Offset((*arg.Pagination.Page - 1) * *arg.Pagination.PerPage)
			}
		}

		if err := tx.Find(&list).Error; err != nil {
			return err
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
