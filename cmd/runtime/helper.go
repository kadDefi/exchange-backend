package runtime

import (
	"context"
	"exchange-backend/internal/config"
	"exchange-backend/internal/pkg/log"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type HandlerArg struct {
	Runtime *Runtime
	Config  *config.Config
}

func NewCommand(use string, handler func(r *Runtime) error) *cobra.Command {
	return &cobra.Command{
		Use: use,
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
			defer cancel()

			logger := log.FromContext(ctx)

			go func() {
				<-ctx.Done()

				logger.Sugar().Info("shutting down...")
			}()
			cfg, err := config.LoadLocal(viper.GetString("config"))
			if err != nil {
				logger.Sugar().Fatal(err)
			}
			r, cleanup, err := build(ctx, cfg)
			if err != nil {
				logger.Sugar().Fatal(err)
			}
			defer cleanup()

			if err := handler(r); err != nil {
				logger.Sugar().Error(err)
			}
		},
	}
}
