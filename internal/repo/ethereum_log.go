package repo

import (
	"context"
	"encoding/json"

	"exchange-backend/internal/domain"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm/clause"
)

func (r *Repo) UpsertTypesLog(ctx context.Context, ms ...types.Log) error {
	arr := make([]*domain.EthereumLog, 0, len(ms))
	for _, v := range ms {
		v.MarshalJSON()
		arr = append(arr, &domain.EthereumLog{
			Address: v.Address.String(),
			Topics: (func() string {
				bs, _ := json.Marshal(v.Topics)
				return string(bs)
			})(),
			Data:        hexutil.Bytes(v.Data).String(),
			BlockNumber: v.BlockNumber,
			TxHash:      v.TxHash.String(),
			TxIndex:     v.TxIndex,
			BlockHash:   v.BlockHash.String(),
			LogIndex:    v.Index,
			Removed:     v.Removed,
		})
	}

	if err := r.db.WithContext(ctx).Table("ethereum_log").Clauses(clause.Insert{
		Modifier: "IGNORE",
	}).CreateInBatches(arr, 100).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) QueryTypesLog(ctx context.Context, arg *domain.QueryTypesLogArg) ([]types.Log, error) {
	tx := r.db.WithContext(ctx).Model(&domain.EthereumLog{})

	if arg.Pagination.PerPage != nil {
		tx = tx.Limit(*arg.Pagination.PerPage)
		if arg.Pagination.Page != nil {
			tx = tx.Offset(*arg.Pagination.Page * *arg.Pagination.PerPage)
		}
	}

	switch arg.Order {
	case domain.QueryTypesLogArgOrderSeqAsc:
		tx = tx.Order("block_number ASC, log_index ASC")
	case domain.QueryTypesLogArgOrderSeqDesc:
		tx = tx.Order("block_number DESC, log_index DESC")
	default:
		tx = tx.Order("id ASC")
	}

	if len(arg.AddressIn) > 0 {
		tx = tx.Where("address IN ?", arg.AddressIn)
	}

	if arg.LogPositionGte != nil {
		tx = tx.Where("block_number > ? OR (block_number = ? AND log_index >= ?)", arg.LogPositionGte.BlockNumber, arg.LogPositionGte.BlockNumber, arg.LogPositionGte.LogIndex)
	}
	if arg.LogPositionLte != nil {
		tx = tx.Where("block_number < ? OR (block_number = ? AND log_index <= ?)", arg.LogPositionLte.BlockNumber, arg.LogPositionLte.BlockNumber, arg.LogPositionLte.LogIndex)
	}

	var list []*domain.EthereumLog
	if err := tx.Find(&list).Error; err != nil {
		return nil, err
	}

	arr := make([]types.Log, 0, len(list))
	for _, v := range list {
		arr = append(arr, v.TypesLog())
	}

	return arr, nil
}
