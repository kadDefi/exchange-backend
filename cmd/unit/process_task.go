package unit

import (
	"exchange-backend/cmd/runtime"
)

func init() {
	Cmd.AddCommand(runtime.NewCommand(
		"process-task",
		func(r *runtime.Runtime) error {
			ctx := r.Ctx

			//r.JobClient.RegisterHandler(domain.TaskRefreshTokenURL{}, r.Service.ProcessTaskRefreshTokenURL)

			go r.JobClient.Pull()

			<-ctx.Done()

			return nil
		},
	))
}
