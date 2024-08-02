package service

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

var (
	ErrNotMatchedLog = errors.New("not matched log")
)

type EventHandler func(context.Context, types.Log) error

type ContractBase struct {
	service *Service

	Address common.Address

	EventHandler map[common.Hash]EventHandler
}

func NewContractBase(address string, service *Service) *ContractBase {
	return &ContractBase{
		Address:      common.HexToAddress(address),
		service:      service,
		EventHandler: make(map[common.Hash]EventHandler),
	}
}

func (c *ContractBase) ProcessLog(ctx context.Context, tl types.Log) (bool, error) {
	if tl.Address.String() != c.Address.String() {
		return false, ErrNotMatchedLog
	}

	if len(tl.Topics) == 0 {
		return false, ErrNotMatchedLog
	}

	if handler, ok := c.EventHandler[tl.Topics[0]]; ok {
		return true, handler(ctx, tl)
	}

	return false, nil
}

func (c *ContractBase) RegisterEventHandler(id common.Hash, handler EventHandler) {
	c.EventHandler[id] = handler
}
