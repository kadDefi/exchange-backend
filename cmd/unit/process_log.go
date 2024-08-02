package unit

import (
	"math"
	"sort"
	"time"

	"exchange-backend/cmd/runtime"
	"exchange-backend/internal/domain"
	"exchange-backend/internal/pkg/helper"
	"exchange-backend/internal/pkg/log"
	"exchange-backend/internal/service"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func init() {
	Cmd.AddCommand(runtime.NewCommand(
		"process-log",
		func(r *runtime.Runtime) error {
			ctx := r.Ctx

			contractServiceMap := make(map[string]service.ContractService)

			one := func() (bool, error) {
				contracts, err := r.Service.QueryContract(ctx, &domain.QueryContractArg{})
				if err != nil {
					return false, err
				}

				if len(contracts) == 0 {
					return false, nil
				}
				for _, v := range contracts {
					cs, err := r.Service.GetContractService(ctx, v.Address, v.Schema)
					if err != nil {
						if err != domain.ErrNotFound {
							return false, err
						}
					} else {
						contractServiceMap[v.Address] = cs
					}
				}

				contracts = lo.Filter(contracts, func(v *domain.Contract, index int) bool {
					// not implemented contract service
					if _, ok := contractServiceMap[v.Address]; !ok {
						return false
					}

					return v.SyncedAtBlockNumber > v.ProcessedAtBlockNumber || (v.SyncedAtBlockNumber == v.ProcessedAtBlockNumber && v.ProcessedAtLogIndex != math.MaxInt32)
				})

				if len(contracts) == 0 {
					return false, nil
				}
				allContractsLength := len(contracts)

				minSyncedContract := lo.MinBy(contracts, func(a *domain.Contract, b *domain.Contract) bool {
					return a.SyncedAtBlockNumber < b.SyncedAtBlockNumber
				})

				sort.Slice(contracts, func(i, j int) bool {
					return contracts[i].ProcessedAtBlockNumber < contracts[j].ProcessedAtBlockNumber || (contracts[i].ProcessedAtBlockNumber == contracts[j].ProcessedAtBlockNumber && contracts[i].ProcessedAtLogIndex < contracts[j].ProcessedAtLogIndex)
				})

				logPositionGte := &domain.LogPosition{
					BlockNumber: contracts[0].ProcessedAtBlockNumber,
					LogIndex:    contracts[0].ProcessedAtLogIndex + 1,
				}

				var logPositionLte *domain.LogPosition
				for _, v := range contracts {
					if v.ProcessedAtBlockNumber != contracts[0].ProcessedAtBlockNumber || v.ProcessedAtLogIndex != contracts[0].ProcessedAtLogIndex {
						logPositionLte = &domain.LogPosition{
							BlockNumber: v.ProcessedAtBlockNumber,
							LogIndex:    v.ProcessedAtLogIndex,
						}
						break
					}
				}

				if logPositionLte == nil || logPositionLte.BlockNumber > minSyncedContract.SyncedAtBlockNumber {
					logPositionLte = &domain.LogPosition{
						BlockNumber: minSyncedContract.SyncedAtBlockNumber,
						LogIndex:    math.MaxInt32,
					}
				}

				contracts = lo.Filter(contracts, func(v *domain.Contract, index int) bool {
					return v.ProcessedAtBlockNumber == contracts[0].ProcessedAtBlockNumber && v.ProcessedAtLogIndex == contracts[0].ProcessedAtLogIndex
				})

				contractMap := lo.SliceToMap(contracts, func(v *domain.Contract) (string, *domain.Contract) {
					return v.Address, v
				})

				sort.Slice(contracts, func(i, j int) bool {
					return contracts[i].ProcessedAtLogIndex < contracts[j].ProcessedAtLogIndex
				})

				if len(contracts) == 0 {
					return false, nil
				}

				addresses := make([]string, 0, len(contracts))
				for _, v := range contracts {
					addresses = append(addresses, v.Address)
				}

				logs, err := r.Service.QueryTypesLog(ctx, &domain.QueryTypesLogArg{
					Pagination: domain.Pagination{
						PerPage: helper.New(100),
					},
					AddressIn:      addresses,
					Order:          domain.QueryTypesLogArgOrderSeqAsc,
					LogPositionGte: logPositionGte,
					LogPositionLte: logPositionLte,
				})
				if err != nil {
					return false, err
				}

				for _, v := range logs {
					log.FromContext(ctx).Sugar().Infof("find log [%s] [%d - %d]", v.Address, v.BlockNumber, v.Index)
					v := v
					contract, ok := contractMap[v.Address.String()]
					if !ok {
						continue
					}
					contractService, ok := contractServiceMap[v.Address.String()]
					if !ok {
						return false, errors.Errorf("contract service %s not found [%s]", contract.Schema, v.Address)
					}

					log.FromContext(ctx).Sugar().Infof("processing log [%s] [%d - %d]", v.Address, v.BlockNumber, v.Index)

					if _, err := contractService.ProcessLog(ctx, v); err != nil {
						return false, err
					}

					if err := r.Service.UpdateContractProcessedAt(ctx, addresses, v.BlockNumber, int(v.Index)); err != nil {
						return false, err
					}
				}

				if len(logs) < 100 {
					if err := r.Service.UpdateContractProcessedAt(ctx, addresses, logPositionLte.BlockNumber, logPositionLte.LogIndex); err != nil {
						return false, err
					}
				}

				return len(logs) == 100 || allContractsLength != len(contracts), nil
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
