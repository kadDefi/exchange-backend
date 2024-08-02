package service

import (
	"context"
	"exchange-backend/internal/domain"
)

func (s *Service) QueryMarket(ctx context.Context, arg *domain.QueryMarketArg) (*domain.MarketResp, error) {
	var (
		resp = new(domain.MarketResp)
	)

	nftItems, total, err := s.repo.QueryMarket(ctx, arg)
	if err != nil {
		return nil, err
	}
	resp.Items = nftItems
	resp.Total = total
	return resp, nil
}
