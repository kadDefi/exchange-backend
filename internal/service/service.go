package service

import (
	"context"
	"exchange-backend/internal/config"
	"exchange-backend/internal/pkg/ethereum"
	"exchange-backend/internal/pkg/etherscan"
	"exchange-backend/internal/pkg/job"
	"exchange-backend/internal/pkg/locker"
	"exchange-backend/internal/repo"
	"github.com/google/wire"
	"github.com/robfig/cron"
)

var ProviderSet = wire.NewSet(
	NewService,
)

type Service struct {
	ctx             context.Context
	cfg             *config.Config
	repo            *repo.Repo
	etherscanClient *etherscan.Client
	ethereumClient  *ethereum.Client
	lockerClient    *locker.Client
	jobClient       *job.Client
	cron            *cron.Cron
}

func NewService(
	ctx context.Context,
	cfg *config.Config,
	repo *repo.Repo,
	etherscanClient *etherscan.Client,
	ethereumClient *ethereum.Client,
	lockerClient *locker.Client,
	jobClient *job.Client,
) *Service {
	return &Service{
		ctx:             ctx,
		cfg:             cfg,
		repo:            repo,
		etherscanClient: etherscanClient,
		ethereumClient:  ethereumClient,
		lockerClient:    lockerClient,
		jobClient:       jobClient,
		cron:            cron.New(),
	}
}
