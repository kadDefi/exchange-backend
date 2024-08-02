package domain

import (
	"encoding/json"
	"time"
)

type ContractSchema string

const (
	ContractSchemaExchange ContractSchema = "exchange"
)

type Contract struct {
	ID        int64     `json:"id,string" gorm:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`

	Address string         `json:"address" gorm:"address"`
	Schema  ContractSchema `json:"schema" gorm:"schema"`
	Name    string         `json:"name" gorm:"name"`

	CreatedAtBlockNumber   uint64 `json:"created_at_block_number" gorm:"created_at_block_number"`
	SyncedAtBlockNumber    uint64 `json:"synced_at_block_number" gorm:"synced_at_block_number"`
	ProcessedAtBlockNumber uint64 `json:"processed_at_block_number" gorm:"processed_at_block_number"`
	ProcessedAtLogIndex    int    `json:"processed_at_log_index" gorm:"processed_at_log_index"`
}

func (m Contract) MarshalBinary() ([]byte, error)     { return json.Marshal(m) }
func (m *Contract) UnmarshalBinary(data []byte) error { return json.Unmarshal(data, m) }

type QueryContractArg struct {
	Pagination Pagination `json:"pagination"`

	AddressIn []string `json:"address_in"`
}
