package repo

import (
	"context"
	"exchange-backend/internal/domain"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func (r *Repo) CreateOrder(ctx context.Context, m *domain.Order) error {
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) QueryOrder(ctx context.Context, arg *domain.QueryOrderArg) ([]*domain.Order, int, error) {
	var (
		list  []*domain.Order
		total int
	)

	tx := r.db.WithContext(ctx).Model(&domain.Order{})

	if arg.SellerAddress != nil {
		tx = tx.Where("seller_address < ?", *arg.SellerAddress)
	}

	if arg.BuyerAddress != nil {
		tx = tx.Where("buyer_address < ?", *arg.BuyerAddress)
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
