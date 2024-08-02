package service

import (
	"context"
	"encoding/json"
	"exchange-backend/internal/domain"
	"exchange-backend/internal/pkg/log"
)

func (s *Service) ProcessTaskRefreshTokenURL(ctx context.Context, arg string) error {
	var t domain.TaskRefreshTokenURL
	if err := json.Unmarshal([]byte(arg), &t); err != nil {
		log.FromContext(ctx).Sugar().Errorf("failed to unmarshal: %v", err)
		return err
	}

	if ms, err := s.repo.GetNFTItemByAddress(ctx, t.CollectionAddress, t.TokenID); err != nil {
		log.FromContext(ctx).Sugar().Infof("NFT item %s@%s not found", t.CollectionAddress, t.TokenID)
		return err
	} else {
		if ms.TokenURI == "" {
			if ExchangeAddrContract, err := NewContractExchange(t.CollectionAddress, s); err != nil {
				log.FromContext(ctx).Sugar().Infof("NewContractExchange Error: %v", err)
				return err
			} else {
				if tokenUrl, err := ExchangeAddrContract.GetTokenURI(ctx, ms.TokenIndex); err != nil {
					log.FromContext(ctx).Sugar().Infof("Get Token URI Error: %v", err)
					return err
				} else {
					ms.TokenURI = tokenUrl
					if err := s.repo.UpdateNFTItem(ctx, ms); err != nil {
						log.FromContext(ctx).Sugar().Infof("Update NFT Item Error: %v", err)
						return err
					}
				}
			}
		}
	}

	return nil
}
