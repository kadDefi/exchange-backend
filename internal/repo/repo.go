package repo

import (
	"context"
	corelog "log"
	"os"
	"sync"
	"time"

	"exchange-backend/internal/config"
	"exchange-backend/internal/pkg/locker"
	"exchange-backend/internal/pkg/log"
	"exchange-backend/internal/pkg/snowflake"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var ProviderSet = wire.NewSet(
	ProvideRepo,
)

type Repo struct {
	snowflakeNode *snowflake.Node
	db            *gorm.DB
	cache         *redis.Client
	locker        *locker.Client
	mutex         sync.Mutex
}

func connectDB(
	ctx context.Context,
	cfg *config.Config,
	snowflakeNode *snowflake.Node,
) (*gorm.DB, error) {
	log.FromContext(ctx).Sugar().Info("[Repo] [MySQL] Connecting")

	errChan := make(chan error, 1)

	var db *gorm.DB
	go func() {
		d, err := gorm.Open(mysql.Open(cfg.Database.DataSourceName), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.New(
				corelog.New(os.Stdout, "\r\n", corelog.LstdFlags),
				logger.Config{
					SlowThreshold:             1000 * time.Millisecond,
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: true,
					Colorful:                  true,
				},
			),
		})

		if err != nil {
			errChan <- err
			return
		}

		db = d
		errChan <- nil
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return nil, errors.Wrap(err, "failed to connect to db")
		}
	case <-time.After(5 * time.Second):
		return nil, errors.Errorf("timeout to connect to db")
	}

	log.FromContext(ctx).Sugar().Info("[Repo] [MySQL] Connected")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	nctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log.FromContext(ctx).Sugar().Info("[Repo] [MySQL] PING database")
	if err := sqlDB.PingContext(nctx); err != nil {
		return nil, err
	}
	log.FromContext(ctx).Sugar().Info("[Repo] [MySQL] PONG database")

	if err := db.Use(NewSnowflakeIDPlugin(snowflakeNode)); err != nil {
		return nil, err
	}

	return db, nil
}

func connectCache(
	ctx context.Context,
	cfg *config.Config,
) (*redis.Client, error) {
	log.FromContext(ctx).Sugar().Info("[Repo] [Cache] Connecting")

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

	log.FromContext(ctx).Sugar().Info("[Repo] [Cache] Connected")

	return cache, nil
}

func ProvideRepo(
	ctx context.Context,
	cfg *config.Config,
	snowflakeNode *snowflake.Node,
	locker *locker.Client,
) (*Repo, func(), error) {
	db, err := connectDB(ctx, cfg, snowflakeNode)
	if err != nil {
		return nil, nil, err
	}

	cache, err := connectCache(ctx, cfg)
	if err != nil {
		return nil, nil, err
	}

	r := &Repo{
		db:            db,
		cache:         cache,
		snowflakeNode: snowflakeNode,
		locker:        locker,
		mutex:         sync.Mutex{},
	}

	return r, r.Stop, nil
}

func (r *Repo) Stop() {
	log.FromContext(context.Background()).Sugar().Info("[Repo] [MySQL] Disconnecting")

	sqlDB, err := r.db.DB()
	if err != nil {
		log.FromContext(context.Background()).Sugar().Errorf("[Repo] [MySQL] failed to get sql db: %v", err)
	} else {
		if err := sqlDB.Close(); err != nil {
			log.FromContext(context.Background()).Sugar().Errorf("[Repo] [MySQL] failed to close sql db: %v", err)
		} else {
			log.FromContext(context.Background()).Sugar().Info("[Repo] [MySQL] Disconnected")
		}
	}

	log.FromContext(context.Background()).Sugar().Info("[Repo] [Cache] Disconnecting")

	if err := r.cache.Close(); err != nil {
		log.FromContext(context.Background()).Sugar().Errorf("[Repo] [Cache] failed to close cache: %v", err)
	} else {
		log.FromContext(context.Background()).Sugar().Info("[Repo] [Cache] Disconnected")
	}
}
