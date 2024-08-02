package domain

import "time"

type Market struct {
	ID        uint64    `json:"id,string" gorm:"column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`

	CollectionAddress string `json:"collection_address" gorm:"column:collection_address"`
	SellerAddress     string `json:"seller_address" gorm:"column:seller_address"`

	TokenID    string  `json:"token_id" gorm:"column:token_id"`
	TokenURI   string  `json:"token_uri" gorm:"column:token_uri"`
	Price      string  `json:"price" gorm:"column:price"`
	PriceFloat float64 `json:"price_float" gorm:"column:price_float"`
	Status     uint8   `json:"status" gorm:"column:status"`
}

func (Market) TableName() string {
	return "market"
}

type QueryMarketArg struct {
	TokenIdIn     []string   `form:"tokenIdIn"`
	Status        *int64     `form:"status,default=1"`
	PriceGt       *float64   `form:"PriceGt"`
	PriceLt       *float64   `form:"PriceLt"`
	SellerAddress *string    `form:"sellerAddress"`
	Pagination    Pagination `form:"pagination"`
}

type MarketResp struct {
	Items []*Market `json:"items"`
	Total int       `json:"total"`
}
