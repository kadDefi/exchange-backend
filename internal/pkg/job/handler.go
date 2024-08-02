package job

import (
	"context"
	"runtime/debug"

	"exchange-backend/internal/pkg/log"
	"github.com/gocraft/work"
	"go.uber.org/zap"
)

type Handler func(ctx context.Context, arg string) error

func (c *Client) RegisterHandler(a interface{}, h Handler) {
	logger := log.FromContext(c.ctx)
	var (
		name string
	)
	type iName interface {
		Name() string
	}

	if i, ok := a.(iName); ok {
		name = i.Name()
	}

	workOptions := work.JobOptions{
		MaxConcurrency: 10,
	}
	type iMaxConcurrency interface {
		MaxConcurrency() uint
	}
	if i, ok := a.(iMaxConcurrency); ok {
		workOptions.MaxConcurrency = i.MaxConcurrency()
	}

	type iPriority interface {
		Priority() uint
	}
	if i, ok := a.(iPriority); ok {
		workOptions.Priority = i.Priority()
	}

	if name != "" && h != nil {
		logger.Sugar().Infof("[JOB] Registering handler for %s", name)
		c.pool.JobWithOptions(name, workOptions, func(j *work.Job) error {
			defer func() {
				if r := recover(); r != nil {
					logger.Sugar().Errorf("stacktrace from panic: \n" + string(debug.Stack()))
				}
			}()

			arg := j.ArgString("arg")
			if err := j.ArgError(); err != nil {
				return err
			}

			ctx := log.WithContext(
				c.ctx,
				log.FromContext(c.ctx).With(
					zap.String("job_id", j.ID),
				),
			)

			if err := h(ctx, arg); err != nil {
				log.FromContext(ctx).Sugar().Errorf("Error handling job_consumer %s: %s", name, err)
				return err
			}

			return nil
		})
	}
}
