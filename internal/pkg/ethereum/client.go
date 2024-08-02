package ethereum

import (
	"context"
	"exchange-backend/internal/config"
	"exchange-backend/internal/pkg/log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gammazero/workerpool"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

var ProviderSet = wire.NewSet(
	ProvideClient,
)

type Client struct {
	Client *ethclient.Client

	cache   *redis.Client
	chainID *big.Int

	workerPool *workerpool.WorkerPool

	cfg *config.Config
}

func ProvideClient(
	ctx context.Context,
	cfg *config.Config,
) (*Client, func(), error) {
	client, err := connectETHClient(ctx, cfg)
	if err != nil {
		return nil, nil, err
	}

	cache, err := connectCache(ctx, cfg)
	if err != nil {
		return nil, nil, err
	}

	c := &Client{
		Client: client,

		cache: cache,

		workerPool: workerpool.New(20),

		cfg: cfg,
	}

	return c, c.Stop, nil
}

func connectETHClient(
	ctx context.Context,
	cfg *config.Config,
) (*ethclient.Client, error) {
	log.FromContext(ctx).Sugar().Info("[Ethereum] [ETHClient] Connecting")

	errChan := make(chan error, 1)

	var client *ethclient.Client
	go func() {
		c, err := ethclient.DialContext(ctx, cfg.Ethereum.RPCURL)
		if err != nil {
			errChan <- err
			return
		}

		client = c
		errChan <- nil
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return nil, err
		}
	case <-time.After(5 * time.Second):
		return nil, errors.Errorf("timeout to connect to eth client")
	}

	log.FromContext(ctx).Sugar().Info("[Ethereum] [ETHClient] Connected")

	return client, nil
}

func connectCache(
	ctx context.Context,
	cfg *config.Config,
) (*redis.Client, error) {
	log.FromContext(ctx).Sugar().Info("[Ethereum] [Cache] Connecting")

	errChan := make(chan error, 1)

	var cache *redis.Client
	go func() {
		c := redis.NewClient(&redis.Options{
			Addr:     cfg.Redis.Cache.Addr,
			Password: cfg.Redis.Cache.Password,
			DB:       cfg.Redis.Cache.DB,
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

	log.FromContext(ctx).Sugar().Info("[Ethereum] [Cache] Connected")

	return cache, nil
}

func (c *Client) Stop() {
	log.FromContext(context.Background()).Sugar().Info("[Ethereum] [ETHClient] Stopping")

	c.Client.Close()

	log.FromContext(context.Background()).Sugar().Info("[Ethereum] [ETHClient] Stopped")

	log.FromContext(context.Background()).Sugar().Info("[Ethereum] [Cache] Stopping")

	if err := c.cache.Close(); err != nil {
		log.FromContext(context.Background()).Sugar().Errorf("[Ethereum] [Cache] Failed to close: %v", err)
	} else {
		log.FromContext(context.Background()).Sugar().Info("[Ethereum] [Cache] Stopped")
	}
}
