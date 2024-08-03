package domain

import "time"

type NFTItem struct {
	ID        uint64    `json:"id,string" gorm:"column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`

	CollectionAddress string `json:"collection_address" gorm:"column:collection_address"`
	OwnerAddress      string `json:"owner_address" gorm:"column:owner_address"`

	TokenIndex  int64  `json:"token_index" gorm:"column:token_index"`
	TokenID     string `json:"token_id" gorm:"column:token_id"`
	MetaData    string `json:"meta_data" gorm:"column:meta_data"`
	TokenURl    string `json:"token_url" gorm:"column:token_url"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	NftPrice    string `json:"nftPrice" gorm:"column:nft_price"`
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

type Metadate struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	NftPrice    string `json:"nftPrice"`
	NftPhone    string `json:"nftPhone"`
	QQNum       string `json:"qqNum"`
	WeChatNum   string `json:"weChatNum"`
	TelegramNum string `json:"telegramNum"`
	SkypeNum    string `json:"skypeNum"`
	BbmNum      string `json:"bbmNum"`
}
