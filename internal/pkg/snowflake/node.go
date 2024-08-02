package snowflake

import (
	"context"
	"fmt"
	"time"

	"exchange-backend/internal/config"
	"exchange-backend/internal/pkg/log"
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
)

var ProviderSet = wire.NewSet(
	NewNode,
)

type Node struct {
	ctx  context.Context
	err  error
	rdb  *redis.Client
	node *snowflake.Node

	nodeId       int
	instanceName string
}

var (
	ErrUninitialized      = errors.New("node is uninitialized")
	ErrInsufficientNodeId = errors.New("insufficient node id")
	ErrNodeIdRegistered   = errors.New("node id is registered")
)

func NewNode(
	ctx context.Context,
	cfg *config.Config,
) (*Node, func(), error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Snowflake.Addr,
		Password: cfg.Redis.Snowflake.Password,
		DB:       cfg.Redis.Snowflake.DB,
	})

	log.FromContext(ctx).Sugar().Info("[SNOWFLAKE] PING redis")
	nctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := rdb.Ping(nctx).Err(); err != nil {
		return nil, nil, err
	}
	log.FromContext(ctx).Sugar().Info("[SNOWFLAKE] PONG redis")

	n := &Node{
		ctx:          ctx,
		err:          ErrUninitialized,
		rdb:          rdb,
		instanceName: ksuid.New().String(),
	}

	log.FromContext(ctx).Sugar().Infof("[SNOWFLAKE] initializing instance name: %s", n.instanceName)
	if err := n.init(ctx); err != nil {
		return nil, nil, err
	}
	log.FromContext(ctx).Sugar().Infof("[SNOWFLAKE] initialized instance name: %s", n.instanceName)

	return n, func() {
		log.FromContext(ctx).Sugar().Info("[SNOWFLAKE] Closing")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := n.deregister(ctx, n.nodeId); err != nil {
			log.FromContext(ctx).Sugar().Errorf("[SNOWFLAKE] deregister failed: %v", err)
		}

		if err := rdb.Close(); err != nil {
			log.FromContext(ctx).Sugar().Errorf("[SNOWFLAKE] error closing redis client: %s", err)
		}

		log.FromContext(ctx).Sugar().Info("[SNOWFLAKE] Closed")
	}, nil
}

func nodeKey(id int) string {
	return fmt.Sprintf("snowflake-node:%d", id)
}

func (n *Node) init(ctx context.Context) error {
	if n.err != ErrUninitialized {
		return n.err
	}

	for id := 0; id < 16; id++ {
		if err := n.register(ctx, id); err == nil {
			log.FromContext(ctx).Sugar().Infof("[SNOWFLAKE] registered node id: %d", id)
			n.nodeId = id

			node, err := snowflake.NewNode(int64(id))
			if err != nil {
				n.err = err
				return err
			}

			n.node = node
			n.err = nil

			go func(id int) {
				for {
					select {
					case <-ctx.Done():
						n.err = n.deregister(ctx, id)
						return
					case <-time.After(10 * time.Second):
						if err := n.register(ctx, id); err != nil {
							n.err = err
						} else {
							n.err = nil
						}
					}
				}
			}(id)

			return nil
		}
	}

	return ErrInsufficientNodeId
}

func (n *Node) register(ctx context.Context, id int) error {
	if err := n.rdb.Watch(n.ctx, func(tx *redis.Tx) error {
		if existed, err := tx.Get(ctx, nodeKey(id)).Result(); err != nil && !errors.Is(err, redis.Nil) {
			return err
		} else if err == nil && existed != n.instanceName {
			return ErrNodeIdRegistered
		}

		if err := tx.Set(n.ctx, nodeKey(id), n.instanceName, 15*time.Second).Err(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (n *Node) deregister(ctx context.Context, id int) error {
	if n.node == nil {
		return nil
	}
	n.node = nil

	if err := n.rdb.Del(ctx, nodeKey(id)).Err(); err != nil {
		return err
	}

	return nil
}

func (n *Node) Generate(ctx context.Context) (snowflake.ID, error) {
	if n.err != nil {
		return 0, n.err
	}

	return n.node.Generate(), nil
}

func (n *Node) BulkGenerate(ctx context.Context, count int) ([]snowflake.ID, error) {
	if n.err != nil {
		return nil, n.err
	}

	ids := make([]snowflake.ID, count)
	for i := 0; i < count; i++ {
		ids[i] = n.node.Generate()
	}

	return ids, nil
}
