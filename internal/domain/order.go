package domain

import "time"

type Order struct {
	ID        uint64    `json:"id,string" gorm:"column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`

	CollectionAddress string `json:"collection_address" gorm:"column:collection_address"`
	SellerAddress     string `json:"seller_address" gorm:"column:seller_address"`
	BuyerAddress      string `json:"buyer_address" gorm:"column:buyer_address"`

	TokenID     string `json:"token_id" gorm:"column:token_id"`
	TokenURI    string `json:"token_uri" gorm:"column:token_url"`
	Price       string `json:"price" gorm:"column:price"`
	TxHash      string `json:"tx_hash" gorm:"tx_hash"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

func (Order) TableName() string {
	return "order"
}

type QueryOrderArg struct {
	SellerAddress *string    `form:"sellerAddress"`
	BuyerAddress  *string    `form:"buyerAddress"`
	Pagination    Pagination `form:"pagination"`
}

type OrderResp struct {
	Items []*Order `json:"items"`
	Total int      `json:"total"`
}
