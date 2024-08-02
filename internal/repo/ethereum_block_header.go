package repo

import (
	"context"
	"encoding/json"

	"exchange-backend/internal/domain"
	"exchange-backend/internal/pkg/helper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm/clause"
)

func (r *Repo) UpsertTypesHeader(ctx context.Context, ms ...*types.Header) error {
	arr := make([]*domain.EthereumBlockHeader, 0, len(ms))
	for _, h := range ms {
		var enc domain.EthereumBlockHeader

		enc.ParentHash = h.ParentHash.String()
		enc.UncleHash = h.UncleHash.String()
		enc.Coinbase = h.Coinbase.String()
		enc.Root = h.Root.String()
		enc.TxHash = h.TxHash.String()
		enc.ReceiptHash = h.ReceiptHash.String()
		enc.Bloom = "0x" + common.Bytes2Hex(h.Bloom.Bytes())
		if h.Difficulty != nil {
			enc.Difficulty = helper.New(h.Difficulty.Uint64())
		}
		if h.Number != nil {
			enc.Number = helper.New(h.Number.Uint64())
		}
		enc.GasLimit = h.GasLimit
		enc.GasUsed = h.GasUsed
		enc.Time = h.Time
		enc.Extra = "0x" + common.Bytes2Hex(h.Extra)
		enc.MixDigest = h.MixDigest.String()
		enc.Nonce = h.Nonce.Uint64()
		if h.BaseFee != nil {
			enc.BaseFee = helper.New(h.BaseFee.Uint64())
		}
		if h.WithdrawalsHash != nil {
			enc.WithdrawalsHash = helper.New(h.WithdrawalsHash.String())
		}
		enc.Hash = h.Hash().String()

		arr = append(arr, &enc)
	}

	if err := r.db.WithContext(ctx).Table("ethereum_block_header").Clauses(clause.Insert{
		Modifier: "IGNORE",
	}).CreateInBatches(arr, 100).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) QueryTypesHeader(ctx context.Context, arg *domain.QueryTypesHeaderArg) ([]*types.Header, error) {
	tx := r.db.WithContext(ctx).Model(&domain.EthereumBlockHeader{})

	if arg.Pagination.PerPage != nil {
		tx = tx.Limit(*arg.Pagination.PerPage)
		if arg.Pagination.Page != nil {
			tx = tx.Offset(*arg.Pagination.Page * *arg.Pagination.PerPage)
		}
	}

	if len(arg.BlockNumberIn) > 0 {
		tx = tx.Where("block_number IN ?", arg.BlockNumberIn)
	}
	if len(arg.HashIn) > 0 {
		tx = tx.Where("hash IN ?", arg.HashIn)
	}

	var list []*domain.EthereumBlockHeader
	if err := tx.Find(&list).Error; err != nil {
		return nil, err
	}

	arr := make([]*types.Header, 0, len(list))
	for _, v := range list {
		var dec types.Header

		bs, _ := json.Marshal(v)
		if err := dec.UnmarshalJSON(bs); err != nil {
			return nil, err
		}
		arr = append(arr, &dec)
	}

	return arr, nil
}
