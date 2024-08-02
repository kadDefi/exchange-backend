package unit

import (
	"math/big"
	"time"

	"exchange-backend/cmd/runtime"
	"exchange-backend/internal/domain"
	"exchange-backend/internal/pkg/ethereum"
	"exchange-backend/internal/service"
	"github.com/ethereum/go-ethereum/common"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"
)

func init() {
	Cmd.AddCommand(runtime.NewCommand(
		"sync-log",
		func(r *runtime.Runtime) error {
			ctx := r.Ctx

			one := func() (bool, error) {
				contracts, err := r.Service.QueryContract(ctx, &domain.QueryContractArg{})
				if err != nil {
					return false, err
				}

				group := lo.GroupBy(contracts, func(v *domain.Contract) uint64 {
					return v.SyncedAtBlockNumber
				})

				currentBlockNumber, err := r.EthereumClient.BlockNumber(ctx)
				if err != nil {
					return false, err
				}
				currentBlockNumber = currentBlockNumber - 1

				goon := false
				eg, ectx := errgroup.WithContext(ctx)
				for fromBlockNumber, contracts := range group {
					if currentBlockNumber <= fromBlockNumber {
						continue
					}

					fromBlockNumber := fromBlockNumber
					contracts := contracts

					eg.Go(func() error {
						toBlockNumber := fromBlockNumber + 2000
						if toBlockNumber > currentBlockNumber {
							toBlockNumber = currentBlockNumber
						} else {
							goon = true
						}

						if err := r.Service.SyncTypesLog(ectx, &service.SyncTypesLogArg{
							FilterQuery: ethereum.FilterQuery{
								Addresses: (func() []common.Address {
									var addresses []common.Address
									for _, contract := range contracts {
										addresses = append(addresses, common.HexToAddress(contract.Address))
									}
									return addresses
								})(),
								FromBlock: big.NewInt(int64(fromBlockNumber)),
								ToBlock:   big.NewInt(int64(toBlockNumber)),
							},
						}); err != nil {
							return err
						}

						return nil
					})
				}

				if err := eg.Wait(); err != nil {
					return false, err
				}

				return goon, nil
			}

			for {
				select {
				case <-ctx.Done():
					return nil
				default:
					goon, err := one()
					if err != nil {
						return err
					}

					if !goon {
						time.Sleep(3 * time.Second)
					}
				}
			}
		},
	))
}
