package locker

import (
	"context"
	"exchange-backend/internal/config"
	"time"

	"exchange-backend/internal/pkg/log"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

var ProviderSet = wire.NewSet(
	ProvideClient,
)

func ProvideClient(
	ctx context.Context,
	cfg *config.Config,
) (*Client, func(), error) {
	rdb, err := connectRdb(ctx, cfg)
	if err != nil {
		return nil, nil, err
	}

	c := &Client{
		rdb:    rdb,
		Client: redislock.New(rdb),
	}

	return c, c.Stop, nil
}

type Client struct {
	*redislock.Client

	rdb *redis.Client
}

func connectRdb(
	ctx context.Context,
	cfg *config.Config,
) (*redis.Client, error) {
	log.FromContext(ctx).Sugar().Info("[Locker] [Rdb] Connecting")

	errChan := make(chan error, 1)

	var cache *redis.Client
	go func() {
		c := redis.NewClient(&redis.Options{
			Addr:     cfg.Redis.Locker.Addr,
			Password: cfg.Redis.Locker.Password,
			DB:       cfg.Redis.Locker.DB,
		})

		errChan <- c.Ping(ctx).Err()
		cache = c
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return nil, errors.Wrap(err, "failed to connect to cache")
		}

	case <-time.After(5 * time.Second):
		return nil, errors.Errorf("timeout to connect to cache")
	}

	log.FromContext(ctx).Sugar().Info("[Locker] [Rdb] Connected")

	return cache, nil
}

func (c *Client) Stop() {
	log.FromContext(context.Background()).Sugar().Info("[Locker] [Rdb] Stopping")

	if err := c.rdb.Close(); err != nil {
		log.FromContext(context.Background()).Sugar().Errorf("[Locker] [Rdb] Failed to close: %s", err)
	} else {
		log.FromContext(context.Background()).Sugar().Info("[Locker] [Rdb] Stopped")
	}
}

func IsErrNotObtained(err error) bool {
	return errors.Is(err, redislock.ErrNotObtained)
}
