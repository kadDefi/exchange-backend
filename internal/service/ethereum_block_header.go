package service

import (
	"context"

	"exchange-backend/internal/domain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (s *Service) UpsertTypesHeader(ctx context.Context, ms ...*types.Header) error {
	return s.repo.UpsertTypesHeader(ctx, ms...)
}

func (s *Service) QueryTypesHeader(ctx context.Context, arg *domain.QueryTypesHeaderArg) ([]*types.Header, error) {
	return s.repo.QueryTypesHeader(ctx, arg)
}

func (s *Service) TypesHeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	list, err := s.QueryTypesHeader(ctx, &domain.QueryTypesHeaderArg{
		HashIn: []string{hash.String()},
	})
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}

	header, err := s.ethereumClient.Client.HeaderByHash(ctx, hash)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get block header by hash %s", hash.Hex())
	}

	if err := s.UpsertTypesHeader(ctx, header); err != nil {
		return nil, err
	}

	return header, nil
}
