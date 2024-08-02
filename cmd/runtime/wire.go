//go:build wireinject
// +build wireinject

//go:generate wire

package runtime

import (
	"context"

	"exchange-backend/internal/api"
	"exchange-backend/internal/config"
	"exchange-backend/internal/pkg/ethereum"
	"exchange-backend/internal/pkg/etherscan"
	"exchange-backend/internal/pkg/job"
	"exchange-backend/internal/pkg/locker"
	"exchange-backend/internal/pkg/snowflake"
	"exchange-backend/internal/repo"
	"exchange-backend/internal/service"
	"github.com/google/wire"
)

func build(
	ctx context.Context,
	cfg *config.Config,
) (*Runtime, func(), error) {

	panic(wire.Build(
		ethereum.ProviderSet,
		etherscan.ProviderSet,
		snowflake.ProviderSet,
		locker.ProviderSet,
		repo.ProviderSet,
		service.ProviderSet,
		job.ProviderSet,
		api.ProviderSet,

		NewRuntime,
	))
}
