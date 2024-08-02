package repo

import (
	"context"

	"exchange-backend/internal/domain"
)

func (r *Repo) GetContractByAddress(ctx context.Context, address string) (*domain.Contract, error) {
	var list []*domain.Contract

	if err := r.db.WithContext(ctx).Where("address = ?", address).Find(&list).Error; err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, domain.ErrNotFound
	}

	return list[0], nil
}

func (r *Repo) CreateContract(ctx context.Context, m *domain.Contract) error {
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdateContractSyncedAtBlockNumber(ctx context.Context, address string, blockNumber uint64) error {
	if err := r.db.WithContext(ctx).Model(&domain.Contract{}).
		Where("address = ?", address).
		Where("synced_at_block_number < ?", blockNumber).
		Update("synced_at_block_number", blockNumber).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdateContractProcessedAt(ctx context.Context, addresses []string, blockNumber uint64, logIndex int) error {
	if err := r.db.WithContext(ctx).Model(&domain.Contract{}).
		Where("address IN ?", addresses).
		Where("(processed_at_block_number < ?) OR (processed_at_block_number = ? AND processed_at_log_index < ?)", blockNumber, blockNumber, logIndex).
		Updates(map[string]interface{}{
			"processed_at_block_number": blockNumber,
			"processed_at_log_index":    logIndex,
		}).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) QueryContract(ctx context.Context, arg *domain.QueryContractArg) ([]*domain.Contract, error) {
	tx := r.db.WithContext(ctx).Model(&domain.Contract{})

	if arg.Pagination.PerPage != nil {
		tx = tx.Limit(*arg.Pagination.PerPage)
		if arg.Pagination.Page != nil {
			tx = tx.Offset((*arg.Pagination.Page - 1) * *arg.Pagination.PerPage)
		}
	}

	if len(arg.AddressIn) > 0 {
		tx = tx.Where("address IN ?", arg.AddressIn)
	}

	var list []*domain.Contract
	if err := tx.Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
