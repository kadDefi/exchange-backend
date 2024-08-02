package domain

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthereumLog struct {
	// ID        int64     `json:"id,string" gorm:"column:id"`
	// CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	// UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	Address     string `json:"address" gorm:"column:address"`
	Topics      string `json:"topics" gorm:"column:topics"`
	Data        string `json:"data" gorm:"column:data"`
	BlockNumber uint64 `json:"block_number" gorm:"column:block_number"`
	TxHash      string `json:"tx_hash" gorm:"column:tx_hash"`
	TxIndex     uint   `json:"tx_index" gorm:"column:tx_index"`
	BlockHash   string `json:"block_hash" gorm:"column:block_hash"`
	LogIndex    uint   `json:"log_index" gorm:"column:log_index"`
	Removed     bool   `json:"removed" gorm:"column:removed"`
}

func (EthereumLog) TableName() string {
	return "ethereum_log"
}

type QueryTypesLogArgOrder string

const (
	QueryTypesLogArgOrderSeqAsc  QueryTypesLogArgOrder = "seq_asc"
	QueryTypesLogArgOrderSeqDesc QueryTypesLogArgOrder = "seq_desc"
)

type LogPosition struct {
	BlockNumber uint64 `json:"block_number" gorm:"block_number"`
	LogIndex    int    `json:"log_index" gorm:"log_index"`
}

type QueryTypesLogArg struct {
	Pagination     Pagination            `json:"pagination"`
	Order          QueryTypesLogArgOrder `json:"order"`
	AddressIn      []string              `json:"address_in"`
	LogPositionGte *LogPosition          `json:"log_position_gt"`
	LogPositionLte *LogPosition          `json:"log_position_lt"`
}

func (m EthereumLog) TypesLog() types.Log {
	return types.Log{
		Address: common.HexToAddress(m.Address),
		Topics: (func() []common.Hash {
			var arr []common.Hash

			json.Unmarshal([]byte(m.Topics), &arr)

			return arr
		})(),
		Data:        hexutil.MustDecode(m.Data),
		BlockNumber: m.BlockNumber,
		TxHash:      common.HexToHash(m.TxHash),
		TxIndex:     m.TxIndex,
		BlockHash:   common.HexToHash(m.BlockHash),
		Index:       m.LogIndex,
		Removed:     m.Removed,
	}
}
