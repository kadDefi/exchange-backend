package job

import (
	"context"

	"exchange-backend/internal/config"
	"exchange-backend/internal/pkg/log"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

type Client struct {
	ctx       context.Context
	redisPool *redis.Pool

	enqueuer *work.Enqueuer
	pool     *work.WorkerPool
}

type Context struct{}

func ProvideClient(
	ctx context.Context,
	cfg *config.Config,
) (*Client, func(), error) {
	redisPool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			opts := make([]redis.DialOption, 0)
			if cfg.Redis.Job.DB != 0 {
				opts = append(opts, redis.DialDatabase(cfg.Redis.Job.DB))
			}
			if cfg.Redis.Job.Password != "" {
				opts = append(opts, redis.DialPassword(cfg.Redis.Job.Password))
			}
			return redis.DialContext(
				ctx,
				"tcp",
				cfg.Redis.Job.Addr,
				opts...,
			)
		},
	}

	enqueuer := work.NewEnqueuer("nft-exchange", redisPool)

	pool := work.NewWorkerPool(Context{}, 5, "nft-exchange", redisPool)

	c := &Client{
		ctx:       ctx,
		redisPool: redisPool,
		enqueuer:  enqueuer,
		pool:      pool,
	}

	c.pool.Middleware(c.logMiddleware)

	return c, c.Stop, nil
}

func (c *Client) Pull() {
	log.FromContext(c.ctx).Sugar().Infof("[JOB] pulling jobs")

	c.pool.Start()
}

func (c *Client) Stop() {
	log.FromContext(c.ctx).Sugar().Infof("[JOB] stopping job consumer")

	c.pool.Stop()

	if err := c.redisPool.Close(); err != nil {
		log.FromContext(c.ctx).Sugar().Errorf("error closing job consumer redis pool: %s", err)
	}
}

func (c *Client) logMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	logger := log.
		FromContext(c.ctx).
		With(zap.String("job_id", job.ID)).
		With(zap.String("job_name", job.Name)).
		With(zap.String("job_arg", job.ArgString("arg")))

	logger.Sugar().Infof("[JOB] processing")

	if err := next(); err != nil {
		logger.Sugar().Errorf("[JOB] failed: %s", err)
		return err
	}

	logger.Sugar().Infof("[JOB] success")
	return nil
}
