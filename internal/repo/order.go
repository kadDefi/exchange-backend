package repo

import (
	"context"
	"exchange-backend/internal/domain"
)

func (r *Repo) CreateOrder(ctx context.Context, m *domain.Order) error {
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}

	return nil
}
