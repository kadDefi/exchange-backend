package api

import (
	"context"
	"exchange-backend/internal/api/middleware"
	"exchange-backend/internal/config"
	"exchange-backend/internal/pkg/log"
	"exchange-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net"
	"net/http"
)

var ProviderSet = wire.NewSet(
	NewBackend,
)

type Backend struct {
	addr string

	server  *http.Server
	service *service.Service
}

func NewBackend(
	cfg *config.Config,
	service *service.Service,
) *Backend {
	e := gin.New()

	e.Use(middleware.Cors())
	e.Use(middleware.Logger())
	e.Use(middleware.Error())
	e.Use(middleware.Recovery())

	b := &Backend{
		addr: cfg.API.Addr,
		server: &http.Server{
			Handler: e,
		},
		service: service,
	}

	e.GET("/user/nft", b.queryUserNFT)
	e.GET("/markets/list", b.queryNFTOrder)
	e.GET("/order/list", b.queryOrderList)
	e.GET("/order/lock", b.queryOrdeLock)

	return b
}

func (b *Backend) Name() string {
	return "Backend"
}

func (b *Backend) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", b.addr)
	if err != nil {
		return err
	}
	log.FromContext(ctx).Sugar().Infof("[API] Listening on %s", lis.Addr().String())

	if err := b.server.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (b *Backend) Stop(ctx context.Context) error {
	if err := b.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
