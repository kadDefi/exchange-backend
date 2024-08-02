package unit

import (
	"exchange-backend/cmd/runtime"
	"exchange-backend/internal/pkg/log"
)

func init() {
	Cmd.AddCommand(runtime.NewCommand(
		"api",
		func(r *runtime.Runtime) error {
			ctx := r.Ctx
			c := r.Backend

			log.FromContext(ctx).Sugar().Info("[API] Starting")
			if err := c.Start(ctx); err != nil {
				return err
			}
			log.FromContext(ctx).Sugar().Info("[API] Running")
			<-ctx.Done()
			if err := c.Stop(ctx); err != nil {
				return err
			}
			log.FromContext(ctx).Sugar().Info("[API] Stopped")
			return nil
		},
	))
}
