package ethereum

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/sync/errgroup"
)

type FilterQuery ethereum.FilterQuery

func (c *Client) FilterLogs(ctx context.Context, q FilterQuery) ([]types.Log, error) {
	var (
		logs []types.Log
		err  error
	)

	c.workerPool.SubmitWait(func() {
		logs, err = c.Client.FilterLogs(
			ctx,
			ethereum.FilterQuery(q),
		)

		if err == nil {
			return
		}

		if !strings.Contains(err.Error(), "query returned more than 2000 results") {
			return
		}

		middle := big.NewInt(0).Div(big.NewInt(0).Add(q.FromBlock, q.ToBlock), big.NewInt(2))

		eg, ctx := errgroup.WithContext(ctx)

		var (
			left  []types.Log
			right []types.Log
		)

		eg.Go(func() error {
			var err error
			left, err = c.FilterLogs(ctx, FilterQuery{
				BlockHash: q.BlockHash,
				FromBlock: q.FromBlock,
				ToBlock:   middle,
				Addresses: q.Addresses,
				Topics:    q.Topics,
			})

			return err
		})

		eg.Go(func() error {
			var err error
			right, err = c.FilterLogs(ctx, FilterQuery{
				BlockHash: q.BlockHash,
				FromBlock: big.NewInt(0).Add(middle, big.NewInt(1)),
				ToBlock:   q.ToBlock,
				Addresses: q.Addresses,
				Topics:    q.Topics,
			})

			return err
		})

		if err = eg.Wait(); err != nil {
			return
		}

		logs = append(left, right...)
	})

	return logs, err
}
