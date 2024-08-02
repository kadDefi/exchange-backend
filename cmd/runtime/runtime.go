package runtime

import (
	"context"
	"exchange-backend/internal/api"
	"exchange-backend/internal/config"
	"exchange-backend/internal/pkg/ethereum"
	"exchange-backend/internal/pkg/job"
	"exchange-backend/internal/pkg/log"

	"exchange-backend/internal/service"
	"go.uber.org/zap"
)

type Runtime struct {
	Ctx            context.Context
	Logger         *zap.Logger
	Cfg            *config.Config
	Service        *service.Service
	EthereumClient *ethereum.Client
	JobClient      *job.Client
	Backend        *api.Backend
}

func NewRuntime(
	ctx context.Context,
	cfg *config.Config,
	service *service.Service,
	ethereumClient *ethereum.Client,
	jobClient *job.Client,
	backend *api.Backend,
) (*Runtime, error) {
	return &Runtime{
		Ctx:            ctx,
		Logger:         log.FromContext(ctx),
		Cfg:            cfg,
		Service:        service,
		EthereumClient: ethereumClient,
		JobClient:      jobClient,
		Backend:        backend,
	}, nil
}
