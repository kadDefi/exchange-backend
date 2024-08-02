package etherscan

import (
	"exchange-backend/internal/config"
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	ProvideClient,
)

type Client struct {
	rc *resty.Client
}

func ProvideClient(cfg *config.Config) *Client {
	return NewClient(
		cfg.Etherscan.APIURL,
		cfg.Etherscan.APIKey,
	)
}

func NewClient(url string, key string) *Client {
	return &Client{
		rc: resty.New().SetBaseURL(url).SetQueryParam("apikey", key),
	}
}
