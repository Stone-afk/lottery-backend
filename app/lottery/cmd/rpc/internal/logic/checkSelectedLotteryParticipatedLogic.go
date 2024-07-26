package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSelectedLotteryParticipatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckSelectedLotteryParticipatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSelectedLotteryParticipatedLogic {
	return &CheckSelectedLotteryParticipatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckSelectedLotteryParticipatedLogic) CheckSelectedLotteryParticipated(in *pb.CheckSelectedLotteryParticipatedReq) (*pb.CheckSelectedLotteryParticipatedResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckSelectedLotteryParticipatedResp{}, nil
}
