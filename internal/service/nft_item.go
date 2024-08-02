package service

import (
	"context"
	"exchange-backend/internal/domain"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *Service) UpsertNFTItem(ctx context.Context, m *domain.NFTItem) error {
	if ms, err := s.repo.GetNFTItemByAddress(ctx, m.CollectionAddress, m.TokenID); err == nil {
		m.ID = ms.ID
		if err := s.repo.UpdateNFTItem(ctx, m); err != nil {
			return errors.Wrapf(err, "failed to update nft item")
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := s.repo.CreateNFTItem(ctx, m); err != nil {
			return errors.Wrapf(err, "failed to create nft item")
		}
	} else {
		return err
	}

	return nil
}
