package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserCreatedLotteryAndTodayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserCreatedLotteryAndTodayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserCreatedLotteryAndTodayLogic {
	return &CheckUserCreatedLotteryAndTodayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserCreatedLotteryAndTodayLogic) CheckUserCreatedLotteryAndToday(in *pb.CheckUserCreatedLotteryAndTodayReq) (*pb.CheckUserCreatedLotteryAndTodayResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckUserCreatedLotteryAndTodayResp{}, nil
}
