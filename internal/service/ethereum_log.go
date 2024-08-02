package service

import (
	"context"
	"math/big"

	"exchange-backend/internal/domain"
	"exchange-backend/internal/pkg/ethereum"
	"exchange-backend/internal/pkg/log"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type SyncTypesLogArg struct {
	ethereum.FilterQuery

	MaxBatchSize uint64
}

func (s *Service) SyncTypesLog(ctx context.Context, arg *SyncTypesLogArg) error {
	if arg.MaxBatchSize == 0 {
		arg.MaxBatchSize = 2000
	}

	if arg.FromBlock == nil {
		arg.FromBlock = big.NewInt(0)
	}
	if arg.ToBlock == nil {
		bn, err := s.ethereumClient.BlockNumber(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to get block number")
		}
		arg.ToBlock = big.NewInt(int64(bn))
	}

	if arg.FromBlock.Cmp(arg.ToBlock) >= 0 {
		return nil
	}

	for i := arg.FromBlock.Uint64(); i < arg.ToBlock.Uint64(); i += arg.MaxBatchSize {
		i := i

		q := ethereum.FilterQuery{
			Addresses: arg.Addresses,
			FromBlock: big.NewInt(int64(i)),
			ToBlock:   big.NewInt(int64(i + arg.MaxBatchSize)),
		}

		if q.ToBlock.Cmp(arg.ToBlock) > 0 {
			q.ToBlock = arg.ToBlock
		}

		logs, err := s.ethereumClient.FilterLogs(ctx, q)
		if err != nil {
			return errors.Wrapf(err, "failed to filter logs from %d - %d for %v", q.FromBlock, q.ToBlock, q.Addresses)
		}

		log.FromContext(ctx).Sugar().Infof("[SyncTypesLog] filter %d logs from %d - %d for %v", len(logs), q.FromBlock, q.ToBlock, q.Addresses)

		if err := s.repo.UpsertTypesLog(ctx, logs...); err != nil {
			return errors.Wrapf(err, "failed to upsert logs from %d - %d for %v", q.FromBlock, q.ToBlock, q.Addresses)
		}

		for _, v := range arg.Addresses {
			if err := s.repo.UpdateContractSyncedAtBlockNumber(ctx, v.String(), q.ToBlock.Uint64()); err != nil {
				return errors.Wrapf(err, "failed to update contract synced_at_block_number from %d - %d for %v", q.FromBlock, q.ToBlock, q.Addresses)
			}
		}
	}

	return nil
}

func (s *Service) QueryTypesLog(ctx context.Context, arg *domain.QueryTypesLogArg) ([]types.Log, error) {
	return s.repo.QueryTypesLog(ctx, arg)
}
