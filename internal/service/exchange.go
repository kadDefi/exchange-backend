package service

import (
	"context"
	"encoding/json"
	"exchange-backend/internal/domain"
	"exchange-backend/internal/pkg/abi"
	"exchange-backend/internal/pkg/helper"
	"exchange-backend/internal/pkg/log"
	"exchange-backend/internal/repo"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"
)

var (
	exchangeABI, _ = abi.ExchangeMetaData.GetAbi()

	exchangeEventTransfer       = newAbiEvent[abi.ExchangeTransfer](exchangeABI.Events["Transfer"])
	exchangeEventSellOrder      = newAbiEvent[abi.ExchangeSellOrder](exchangeABI.Events["SellOrder"])
	exchangeEventMetadataUpdate = newAbiEvent[abi.ExchangeMetadataUpdate](exchangeABI.Events["MetadataUpdate"])
	exchangeEventCancelOrder    = newAbiEvent[abi.ExchangeCancelOrder](exchangeABI.Events["CancelOrder"])
	exchangeEventPurchasedOrder = newAbiEvent[abi.ExchangeTokenPurchased](exchangeABI.Events["TokenPurchased"])
)

type ContractExchange struct {
	*ContractBase
}

func NewContractExchange(
	address string,
	service *Service,
) (*ContractExchange, error) {
	c := &ContractExchange{
		ContractBase: NewContractBase(address, service),
	}

	c.RegisterEventHandler(exchangeEventTransfer.ID, c.processLogTransfer)
	c.RegisterEventHandler(exchangeEventSellOrder.ID, c.processLogSellOrder)
	c.RegisterEventHandler(exchangeEventMetadataUpdate.ID, c.processLogMetadataUpdate)
	c.RegisterEventHandler(exchangeEventCancelOrder.ID, c.processLogCancelOrder)
	c.RegisterEventHandler(exchangeEventCancelOrder.ID, c.processLogCancelOrder)

	return c, nil
}

func (c *ContractExchange) processLogTransfer(ctx context.Context, tl types.Log) error {
	e, err := exchangeEventTransfer.UnpackLog(tl)
	if err != nil {
		return err
	}

	log.FromContext(ctx).Sugar().Infof("[%d] [%s:%s] transfer %s -> %s", tl.BlockNumber, c.Address.String(), e.TokenId.String(), e.From.String(), e.To.String())

	nftItem := domain.NFTItem{
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		CollectionAddress: c.Address.String(),
		OwnerAddress:      e.To.String(),
		TokenIndex:        e.TokenId.Int64(),
		TokenID:           e.TokenId.String(),
	}

	if err := c.service.repo.CreateNFTItem(ctx, &nftItem); err != nil {
		return errors.Wrapf(err, "failed to create NFT item")
	}

	return nil
}

func (c *ContractExchange) processLogMetadataUpdate(ctx context.Context, tl types.Log) error {
	e, err := exchangeEventMetadataUpdate.UnpackLog(tl)
	if err != nil {
		return err
	}

	log.FromContext(ctx).Sugar().Infof("[%d] [%s:%s] MetadataUpdate", tl.BlockNumber, c.Address.String(), e.TokenId.String())

	if ms, err := c.service.repo.GetNFTItemByAddress(ctx, c.Address.String(), e.TokenId.String()); err != nil {
		log.FromContext(ctx).Sugar().Infof("NFT item %s@%s not found", c.Address.String(), e.TokenId.String())
		return err
	} else {
		if ms.TokenURl == "" {
			if tokenUrl, err := c.GetTokenURI(ctx, ms.TokenIndex); err != nil {
				log.FromContext(ctx).Sugar().Infof("Get Token URI Error: %v", err)
				return err
			} else {
				metadata := tokenUrl
				req, err := http.NewRequest("GET", metadata, nil)
				if err != nil {
					return err
				}

				req.Header.Add("Accept", "application/json")
				req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")

				res, err := http.DefaultClient.Do(req)

				defer func() {
					if res != nil {
						_ = res.Body.Close()
					}
				}()
				if res == nil || err != nil {
					return err
				}
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					return err
				}

				data := domain.Metadate{}
				err = json.Unmarshal(body, &data)
				if err != nil {
					log.FromContext(ctx).Sugar().Errorf("Unmarshal Error: %v", err)
					return err
				}

				ms.MetaData = tokenUrl
				ms.TokenURl = data.Image
				if err := c.service.repo.UpdateNFTItem(ctx, ms); err != nil {
					log.FromContext(ctx).Sugar().Infof("Update NFT Item Error: %v", err)
					return err
				}
			}
		}
	}

	return nil
}

func (c *ContractExchange) processLogSellOrder(ctx context.Context, tl types.Log) error {
	e, err := exchangeEventSellOrder.UnpackLog(tl)
	if err != nil {
		return err
	}

	log.FromContext(ctx).Sugar().Infof("[%d] [%s:%s] TokenTransferredToExchange %s -> %s", tl.BlockNumber, c.Address.String(), e.TokenId.String(), e.From.String(), e.Price.String())

	nftItem, err := c.service.repo.GetNFTItemByAddress(ctx, c.Address.String(), e.TokenId.String())
	if err != nil {
		return err
	}

	market := domain.Market{
		CollectionAddress: c.Address.String(),
		TokenID:           e.TokenId.String(),
		Price:             e.Price.String(),
		SellerAddress:     e.From.String(),
		TokenURI:          nftItem.TokenURl,
		PriceFloat:        helper.TokenAmountStrToFloat64(e.Price.String()),
		Status:            repo.ONMARKET,
	}

	if err := c.service.repo.CreateMarket(ctx, &market); err != nil {
		return errors.Wrapf(err, "failed to create market")
	}

	return err
}

func (c *ContractExchange) processLogCancelOrder(ctx context.Context, tl types.Log) error {
	e, err := exchangeEventCancelOrder.UnpackLog(tl)
	if err != nil {
		return err
	}

	log.FromContext(ctx).Sugar().Infof("[%d] [%s:%s] CancelOrder %s", tl.BlockNumber, c.Address.String(), e.TokenId.String())

	if market, err := c.service.repo.GetMarketByAddress(ctx, c.Address.String(), e.TokenId.String(), repo.ONMARKET); err != nil {
		return errors.Wrapf(err, "failed to create market")
	} else {
		market.Status = repo.CANCLE
		if err := c.service.repo.UpdateMarket(ctx, market); err != nil {
			return errors.Wrapf(err, "failed to update market")
		}
	}

	return err
}

func (c *ContractExchange) processLogPurchasedOrder(ctx context.Context, tl types.Log) error {
	e, err := exchangeEventPurchasedOrder.UnpackLog(tl)
	if err != nil {
		return err
	}

	log.FromContext(ctx).Sugar().Infof("[%d] [%s:%s] PurchasedOrder %s =>%s", tl.BlockNumber, c.Address.String(), e.TokenId.String(), e.Seller.String(), e.Buyer.String())

	if market, err := c.service.repo.GetMarketByAddress(ctx, c.Address.String(), e.TokenId.String(), repo.ONMARKET); err != nil {
		return errors.Wrapf(err, "failed to create market")
	} else {
		market.Status = repo.SELLED
		if err := c.service.repo.UpdateMarket(ctx, market); err != nil {
			return errors.Wrapf(err, "failed to update market")
		}

		order := domain.Order{
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			CollectionAddress: c.Address.String(),
			SellerAddress:     e.Seller.String(),
			BuyerAddress:      e.Buyer.String(),
			TokenID:           e.TokenId.String(),
			TokenURI:          market.TokenURI,
			Price:             market.Price,
			TxHash:            e.Raw.TxHash.String(),
		}

		if err := c.service.repo.CreateOrder(ctx, &order); err != nil {
			return errors.Wrapf(err, "failed to create order")
		}
	}

	return err
}

func (c *ContractExchange) GetTokenURI(ctx context.Context, tokenId int64) (string, error) {
	caller, err := abi.NewExchange(c.Address, c.service.ethereumClient.Client)
	if err != nil {
		return "", err
	}

	return caller.TokenURI(nil, new(big.Int).SetInt64(tokenId))
}
