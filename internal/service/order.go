package service

import (
	"context"
	"exchange-backend/internal/domain"
)

func (s *Service) QueryOrder(ctx context.Context, arg *domain.QueryOrderArg) (*domain.OrderResp, error) {
	var (
		resp = new(domain.OrderResp)
	)

	nftItems, total, err := s.repo.QueryOrder(ctx, arg)
	if err != nil {
		return nil, err
	}
	resp.Items = nftItems
	resp.Total = total
	return resp, nil
}
