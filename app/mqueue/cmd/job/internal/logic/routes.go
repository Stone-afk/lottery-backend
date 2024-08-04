package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"looklook/app/mqueue/cmd/job/internal/svc"
	"looklook/app/mqueue/cmd/job/jobtype"
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
	mux := asynq.NewServeMux()

	//scheduler job (looklook)
	//mux.Handle(jobtype.ScheduleSettleRecord, NewSettleRecordHandler(l.svcCtx))

	//defer job (looklook)
	//mux.Handle(jobtype.DeferCloseHomestayOrder, NewCloseHomestayOrderHandler(l.svcCtx))

	//queue job (lottery)
	mux.Handle(jobtype.MsgWxMiniProgramNotifyUser, NewWxMiniProgramNotifyUserHandler(l.svcCtx))

	//schedule job
	mux.Handle(jobtype.ScheduleLotteryDraw, NewLotteryDrawHandler(l.svcCtx))
	//mux.Handle(jobtype.ScheduleWishCheckin, NewWishCheckinHandler(l.svcCtx))
	return mux
}
