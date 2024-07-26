package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserCreatedLotteryAndThisWeekLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserCreatedLotteryAndThisWeekLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserCreatedLotteryAndThisWeekLogic {
	return &CheckUserCreatedLotteryAndThisWeekLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserCreatedLotteryAndThisWeekLogic) CheckUserCreatedLotteryAndThisWeek(in *pb.CheckUserCreatedLotteryAndThisWeekReq) (*pb.CheckUserCreatedLotteryAndThisWeekResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckUserCreatedLotteryAndThisWeekResp{}, nil
}
