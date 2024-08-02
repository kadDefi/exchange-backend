package script

import (
	"exchange-backend/cmd/runtime"
	"exchange-backend/internal/config"
	"exchange-backend/internal/domain"
	"exchange-backend/internal/pkg/log"
	"reflect"
)

func init() {
	Cmd.AddCommand(runtime.NewCommand(
		"init_contract",
		func(r *runtime.Runtime) error {
			ctx := r.Ctx
			logger := log.FromContext(ctx)

			v := reflect.ValueOf(r.Cfg.Contract)
			typeOfS := v.Type()

			for i := 0; i < v.NumField(); i++ {
				name := typeOfS.Field(i).Name
				value := v.Field(i).Interface().(config.Contract)

				if value.Address == "" {
					continue
				}

				if value.SkipInit {
					continue
				}

				logger.Sugar().Infof("init st nft: [%s] %s", name, value.Address)

				// init st nft
				if err := r.Service.InitContract(ctx, &domain.Contract{
					Address:                value.Address,
					Schema:                 domain.ContractSchema(value.Schema),
					Name:                   name,
					CreatedAtBlockNumber:   value.CreatedAtBlockNumber,
					SyncedAtBlockNumber:    0,
					ProcessedAtBlockNumber: 0,
					ProcessedAtLogIndex:    0,
				}); err != nil {
					return err
				}
			}
			return nil
		},
	))
}
