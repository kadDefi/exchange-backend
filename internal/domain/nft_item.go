package domain

import "time"

type NFTItem struct {
	ID        uint64    `json:"id,string" gorm:"column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`

	CollectionAddress string `json:"collection_address" gorm:"column:collection_address"`
	OwnerAddress      string `json:"owner_address" gorm:"column:owner_address"`

	TokenIndex int64  `json:"token_index" gorm:"column:token_index"`
	TokenID    string `json:"token_id" gorm:"column:token_id"`
	TokenURI   string `json:"token_uri" gorm:"column:token_uri"`
}

func (NFTItem) TableName() string {
	return "nft_item"
}

type QueryNFTItemArg struct {
	CollectionAddressIn []string   `form:"collectionAddressIn"`
	TokenIDIn           []string   `form:"tokenIdIn"`
	OwnerAddress        *string    `form:"ownerAddress"`
	Pagination          Pagination `form:"pagination"`
}

type NFTItemResp struct {
	Items []*NFTItem `json:"items"`
	Total int        `json:"total"`
}
