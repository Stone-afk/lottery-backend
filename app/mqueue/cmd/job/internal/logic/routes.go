package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"looklook/app/mqueue/cmd/job/internal/svc"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register register job
func (l *CronJob) Register() *asynq.ServeMux {
	panic("")
}
