package domain

import (
	"fmt"
)

type TaskRefreshTokenURL struct {
	CollectionAddress string `json:"collection_address"`
	TokenID           string `json:"token_id"`
}

func (t TaskRefreshTokenURL) Name() string { return "refresh_token_url" }

func (t TaskRefreshTokenURL) UniqueKey() string {
	return fmt.Sprintf("token_url:%s", t.TokenID)
}
