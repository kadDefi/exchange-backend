package domain

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthereumBlockHeader struct {
	Hash string `json:"hash" gorm:"column:hash"`

	ParentHash  string `json:"parentHash" gorm:"column:parent_hash"`
	UncleHash   string `json:"sha3Uncles" gorm:"column:sha3_uncles"`
	Coinbase    string `json:"miner" gorm:"column:miner"`
	Root        string `json:"stateRoot" gorm:"column:state_root"`
	TxHash      string `json:"transactionsRoot" gorm:"column:transactions_root"`
	ReceiptHash string `json:"receiptsRoot" gorm:"column:receipts_root"`
	Bloom       string `json:"logsBloom" gorm:"column:logs_bloom"`

	Difficulty *uint64 `json:"difficulty,string" gorm:"column:difficulty"`
	Number     *uint64 `json:"number,string" gorm:"column:number"`
	GasLimit   uint64  `json:"gasLimit,string" gorm:"column:gas_limit"`
	GasUsed    uint64  `json:"gasUsed,string" gorm:"column:gas_used"`
	Time       uint64  `json:"timestamp,string" gorm:"column:timestamp"`
	// Difficulty *uint64 `json:"-" gorm:"column:difficulty"`
	// Number     *uint64 `json:"-" gorm:"column:number"`
	// GasLimit   uint64  `json:"-" gorm:"column:gas_limit"`
	// GasUsed    uint64  `json:"-" gorm:"column:gas_used"`
	// Time       uint64  `json:"-" gorm:"column:timestamp"`

	Extra     string `json:"extraData" gorm:"column:extra_data"`
	MixDigest string `json:"mixHash" gorm:"column:mix_hash"`

	Nonce   uint64  `json:"nonce,string" gorm:"column:nonce"`
	BaseFee *uint64 `json:"baseFeePerGas,string" gorm:"column:base_fee_per_gas"`
	// Nonce   uint64  `json:"-" gorm:"column:nonce"`
	// BaseFee *uint64 `json:"-" gorm:"column:base_fee_per_gas"`

	WithdrawalsHash *string `json:"withdrawalsRoot" gorm:"column:withdrawals_root"`
}

func (m EthereumBlockHeader) MarshalJSON() ([]byte, error) {
	uint64ToHexutilBig := func(u *uint64) *hexutil.Big {
		if u == nil {
			return nil
		}

		return (*hexutil.Big)(big.NewInt(int64(*u)))
	}

	uint64ToHexutilUint64 := func(u uint64) hexutil.Uint64 {
		return hexutil.Uint64(u)
	}

	return json.Marshal(map[string]interface{}{
		"hash":             m.Hash,
		"parentHash":       m.ParentHash,
		"sha3Uncles":       m.UncleHash,
		"miner":            m.Coinbase,
		"stateRoot":        m.Root,
		"transactionsRoot": m.TxHash,
		"receiptsRoot":     m.ReceiptHash,
		"logsBloom":        m.Bloom,
		"difficulty":       uint64ToHexutilBig(m.Difficulty),
		"number":           uint64ToHexutilBig(m.Number),
		"gasLimit":         uint64ToHexutilUint64(m.GasLimit),
		"gasUsed":          uint64ToHexutilUint64(m.GasUsed),
		"timestamp":        uint64ToHexutilUint64(m.Time),
		"extraData":        m.Extra,
		"mixHash":          m.MixDigest,
		"nonce":            types.EncodeNonce(m.Nonce),
		"baseFeePerGas":    uint64ToHexutilBig(m.BaseFee),
		"withdrawalsRoot":  m.WithdrawalsHash,
	})
}

func (EthereumBlockHeader) TableName() string {
	return "ethereum_block_header"
}

type QueryTypesHeaderArg struct {
	Pagination    Pagination `json:"pagination"`
	BlockNumberIn []uint64   `json:"block_number_in"`
	HashIn        []string   `json:"hash_in"`
}
