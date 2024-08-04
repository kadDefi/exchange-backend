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
		return err
	}

	log.FromContext(ctx).Sugar().Infof("Processing task refresh token,CollectionAddress: %v, TokenId: %v", t.CollectionAddress, t.TokenID)

	return nil
}
