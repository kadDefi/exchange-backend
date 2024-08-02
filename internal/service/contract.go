package service

import (
	"context"
	"exchange-backend/internal/pkg/etherscan"
	"reflect"

	"exchange-backend/internal/domain"
	"exchange-backend/internal/pkg/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type ContractAssetMetadata struct {
	BlockNumber uint64
	Name        string
	Symbol      string
	TotalSupply string
	Decimals    uint8
}

type ContractService interface {
	ProcessLog(ctx context.Context, tl types.Log) (bool, error)
}

func (s *Service) GetContractService(
	ctx context.Context,
	address string,
	schema domain.ContractSchema,
) (ContractService, error) {
	switch schema {
	case domain.ContractSchemaExchange:
		return NewContractExchange(address, s)
	default:
		return nil, domain.ErrNotFound
	}
}

type abiEvent[T any] abi.Event

func newAbiEvent[T any](event abi.Event) abiEvent[T] {
	if event.ID.String() == ethereum.EmptyHash.String() {
		panic("abi: event ID is nil")
	}

	return abiEvent[T](event)
}

func (e abiEvent[T]) UnpackLog(log types.Log) (out *T, err error) {
	out = reflect.New(reflect.TypeOf(out).Elem()).Interface().(*T)

	if log.Topics[0] != e.ID {
		return out, errors.Errorf("event signature mismatch")
	}

	if len(log.Data) > 0 {
		unpacked, err := e.Inputs.Unpack(log.Data)
		if err != nil {
			return out, err
		}
		if err := e.Inputs.Copy(out, unpacked); err != nil {
			return out, err
		}
	}

	var indexed abi.Arguments
	for _, arg := range e.Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return out, err
	}

	return out, nil
}

func TokenKey(collectionAddress string, tokenID string) string {
	return collectionAddress + ":" + tokenID
}

func (s *Service) QueryContract(ctx context.Context, arg *domain.QueryContractArg) ([]*domain.Contract, error) {
	return s.repo.QueryContract(ctx, arg)
}

func (s *Service) UpdateContractProcessedAt(ctx context.Context, addresses []string, blockNumber uint64, logIndex int) error {
	return s.repo.UpdateContractProcessedAt(ctx, addresses, blockNumber, logIndex)
}

func (s *Service) InitContract(ctx context.Context, m *domain.Contract) error {
	_, err := s.repo.GetContractByAddress(ctx, m.Address)
	if err == nil {
		return nil
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return err
	}

	if !common.IsHexAddress(m.Address) {
		return errors.Errorf("invalid contract %s with address: %s", m.Name, m.Address)
	}
	m.Address = common.HexToAddress(m.Address).String()

	if m.CreatedAtBlockNumber == 0 {
		result, err := s.etherscanClient.GetContractCreation(ctx, &etherscan.GetContractCreationArg{
			ContractAddresses: []string{m.Address},
		})
		if err != nil {
			return errors.Wrapf(err, "failed to get contract creation")
		}
		if len(result) == 0 {
			return errors.New("no result")
		}

		receipt, err := s.ethereumClient.TransactionReceipt(ctx, common.HexToHash(result[0].TxHash))
		if err != nil {
			return err
		}

		m.CreatedAtBlockNumber = receipt.BlockNumber.Uint64()
	}

	// save
	if err := s.repo.CreateContract(ctx, &domain.Contract{
		Address:                m.Address,
		Schema:                 domain.ContractSchema(m.Schema),
		Name:                   m.Name,
		CreatedAtBlockNumber:   m.CreatedAtBlockNumber,
		SyncedAtBlockNumber:    m.CreatedAtBlockNumber - 1,
		ProcessedAtBlockNumber: 0,
		ProcessedAtLogIndex:    0,
	}); err != nil {
		return err
	}

	return nil
}
