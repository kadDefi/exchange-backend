package service

import (
	"context"
	"exchange-backend/internal/domain"
)

func (s *Service) QueryUserNFT(ctx context.Context, arg *domain.QueryNFTItemArg) (*domain.NFTItemResp, error) {
	var (
		resp = new(domain.NFTItemResp)
	)

	nftItems, total, err := s.repo.QueryNFTItems(ctx, arg)
	if err != nil {
		return nil, err
	}
	resp.Items = nftItems
	resp.Total = total
	return resp, nil
}
