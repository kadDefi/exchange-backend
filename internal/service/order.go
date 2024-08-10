package service

import (
	"context"
	"exchange-backend/internal/domain"
	"exchange-backend/internal/pkg/log"
	"github.com/bsm/redislock"
	"time"
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

func (s *Service) QueryOrderLock(ctx context.Context) error {
	log.FromContext(ctx).Sugar().Infof("QueryOrderLock Start!")

	lock, err := s.lockerClient.Obtain(ctx, "OrderLock", 5*time.Second, &redislock.Options{
		RetryStrategy: redislock.LinearBackoff(500 * time.Millisecond),
	})

	if err != nil {
		return err
	}
	defer lock.Release(ctx)

	time.Sleep(5 * time.Second)
	log.FromContext(ctx).Sugar().Infof("========QueryOrderLock End!========")

	return nil
}
