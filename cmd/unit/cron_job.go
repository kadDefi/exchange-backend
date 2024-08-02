package unit

import (
	"exchange-backend/cmd/runtime"
	"github.com/robfig/cron"
)

func init() {
	Cmd.AddCommand(runtime.NewCommand(
		"cron-job",
		func(r *runtime.Runtime) error {
			ctx := r.Ctx
			c := cron.New()

			//c.AddFunc("@every 10m", r.Service.CronMatchNftAndClaimRewards)
			//c.AddFunc("@every 30m", r.Service.CronMatchNftAndStakeApecoin)

			c.Start()
			<-ctx.Done()
			c.Stop()
			return nil
		},
	))
}
