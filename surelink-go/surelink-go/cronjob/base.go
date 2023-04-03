package cronjob

import "context"

type BaseCronJob interface {
	Run(ctx context.Context)
}
