package domain

import (
	"fmt"
)

type TaskRefreshTokenURL struct {
	CollectionAddress string       `json:"collection_address"`
	TokenID           string       `json:"token_id"`
	EthereumLog       *EthereumLog `json:"ethereum_log"`
}

func (t TaskRefreshTokenURL) Name() string { return "refresh_token_url" }

func (t TaskRefreshTokenURL) UniqueKey() string {
	return fmt.Sprintf("token_url:%d%d", t.EthereumLog.BlockNumber, t.EthereumLog.LogIndex)
}
