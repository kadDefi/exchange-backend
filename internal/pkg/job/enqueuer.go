package job

import (
	"context"
	"encoding/json"
	"time"

	"exchange-backend/internal/pkg/log"
	"github.com/gocraft/work"
)

type jobEnqueueOptions struct {
	name      string
	arg       interface{}
	uniqueKey string
	delay     time.Duration
}

type EnqueueOption func(*jobEnqueueOptions)

func WithJobName(name string) EnqueueOption {
	return func(o *jobEnqueueOptions) {
		o.name = name
	}
}

func WithJobArg(arg interface{}) EnqueueOption {
	return func(o *jobEnqueueOptions) {
		o.arg = arg
	}
}

func WithJobUniqueKey(uniqueKey string) EnqueueOption {
	return func(o *jobEnqueueOptions) {
		o.uniqueKey = uniqueKey
	}
}

func WithJobDelay(delay time.Duration) EnqueueOption {
	return func(o *jobEnqueueOptions) {
		o.delay = delay
	}
}

func (c *Client) enqueue(
	ctx context.Context,
	opts ...EnqueueOption,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	options := &jobEnqueueOptions{}
	for _, o := range opts {
		o(options)
	}

	argBs, err := json.Marshal(options.arg)
	if err != nil {
		return err
	}

	var jobID string

	if options.uniqueKey != "" {
		if options.delay != 0 {
			var j *work.ScheduledJob
			j, err = c.enqueuer.EnqueueUniqueIn(
				options.name,
				int64(options.delay.Seconds()),
				work.Q{
					"object_id_": options.uniqueKey,
					"arg":        string(argBs),
				},
			)
			if j != nil {
				jobID = j.ID
			}

		} else {
			var j *work.Job
			j, err = c.enqueuer.EnqueueUnique(
				options.name,
				work.Q{
					"object_id_": options.uniqueKey,
					"arg":        string(argBs),
				},
			)
			if j != nil {
				jobID = j.ID
			}
		}

	} else {
		if options.delay != 0 {
			var j *work.ScheduledJob

			j, err = c.enqueuer.EnqueueIn(
				options.name,
				int64(options.delay.Seconds()),
				work.Q{
					"arg": string(argBs),
				},
			)

			if j != nil {
				jobID = j.ID
			}
		} else {
			var j *work.Job
			j, err = c.enqueuer.Enqueue(
				options.name,
				work.Q{
					"arg": string(argBs),
				},
			)
			if j != nil {
				jobID = j.ID
			}
		}
	}

	if err != nil {
		return err
	}

	if jobID != "" {
		log.FromContext(ctx).Sugar().Infof("[JOB] Enqueued job_consumer %s", jobID)
	}

	return nil
}

func (c *Client) Enqueue(ctx context.Context, a interface{}, opt ...EnqueueOption) error {
	opts := make([]EnqueueOption, 0)
	type iName interface {
		Name() string
	}

	if v, ok := a.(iName); ok {
		opts = append(opts, WithJobName(v.Name()))
	}

	type iUniqueKey interface {
		UniqueKey() string
	}

	if v, ok := a.(iUniqueKey); ok {
		opts = append(opts, WithJobUniqueKey(v.UniqueKey()))
	}

	type iDelay interface {
		Delay() time.Duration
	}

	if v, ok := a.(iDelay); ok {
		opts = append(opts, WithJobDelay(v.Delay()))
	}

	opts = append(opts, WithJobArg(a))
	opts = append(opts, opt...)

	return c.enqueue(ctx, opts...)
}
