package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserCreatedLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserCreatedLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserCreatedLotteryLogic {
	return &CheckUserCreatedLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserCreatedLotteryLogic) CheckUserCreatedLottery(in *pb.CheckUserCreatedLotteryReq) (*pb.CheckUserCreatedLotteryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckUserCreatedLotteryResp{}, nil
}
