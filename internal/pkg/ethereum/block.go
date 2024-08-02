package ethereum

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

func (c *Client) BlockNumber(ctx context.Context) (uint64, error) {
	var (
		bn  uint64
		err error
	)

	cacheKey := "ethereum:block_number"

	if err := c.cache.Get(ctx, cacheKey).Scan(&bn); err != nil {
		if !errors.Is(err, redis.Nil) {
			return 0, err
		}
	}

	bn, err = c.Client.BlockNumber(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get block number")
	}

	if err := c.cache.Set(ctx, cacheKey, bn, 5*time.Second).Err(); err != nil {
		return 0, err
	}

	return bn, nil
}

type typesHeader types.Header

func (h *typesHeader) MarshalBinary() ([]byte, error)    { return rlp.EncodeToBytes(h) }
func (h *typesHeader) UnmarshalBinary(data []byte) error { return rlp.DecodeBytes(data, h) }

func (c *Client) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	var m = new(typesHeader)

	cacheKey := "ethereum:block_header:" + hash.Hex()

	if err := c.cache.Get(ctx, cacheKey).Scan(m); err != nil {
		if !errors.Is(err, redis.Nil) {
			return nil, err
		}
	} else {
		return (*types.Header)(m), nil
	}

	header, err := c.Client.HeaderByHash(ctx, hash)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get block header by hash %s", hash.Hex())
	}

	if err := c.cache.Set(ctx, cacheKey, (*typesHeader)(header), 0).Err(); err != nil {
		return nil, err
	}

	return header, nil
}

func (c *Client) TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	var m = new(types.Receipt)

	cacheKey := "ethereum:transaction_receipt:" + hash.Hex()

	if err := c.cache.Get(ctx, cacheKey).Scan(m); err != nil {
		if !errors.Is(err, redis.Nil) {
			return nil, err
		}
	}

	receipt, err := c.Client.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get transaction receipt by hash %s", hash.Hex())
	}

	if err := c.cache.Set(ctx, cacheKey, receipt, 0).Err(); err != nil {
		return nil, err
	}

	return receipt, nil
}
