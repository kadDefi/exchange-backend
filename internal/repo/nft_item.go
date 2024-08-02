package repo

import (
	"context"
	"exchange-backend/internal/domain"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func (r *Repo) GetNFTItemByAddress(ctx context.Context, collectionAddress string, tokenID string) (*domain.NFTItem, error) {
	nftItem := new(domain.NFTItem)
	if err := r.db.WithContext(ctx).
		Where("collection_address = ?", collectionAddress).
		Where("token_id = ?", tokenID).
		First(nftItem).Error; err != nil {
		return nil, err
	}

	return nftItem, nil
}

func (r *Repo) UpdateNFTItem(ctx context.Context, ms *domain.NFTItem) error {
	if err := r.db.WithContext(ctx).Model(ms).
		Where("id = ?", ms.ID).
		Updates(ms).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) CreateNFTItem(ctx context.Context, m *domain.NFTItem) error {
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) QueryNFTItems(ctx context.Context, arg *domain.QueryNFTItemArg) ([]*domain.NFTItem, int, error) {
	var (
		list  []*domain.NFTItem
		total int
	)

	tx := r.db.WithContext(ctx).Model(&domain.NFTItem{})

	if len(arg.CollectionAddressIn) > 0 {
		tx = tx.Where("collection_address IN ?", arg.CollectionAddressIn)
	}

	if len(arg.TokenIDIn) > 0 {
		tx = tx.Where("token_id IN ?", arg.TokenIDIn)
	}

	if arg.OwnerAddress != nil {
		tx = tx.Where("owner_address = ?", *arg.OwnerAddress)
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
