package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetIsSelectedLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetIsSelectedLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetIsSelectedLotteryLogic {
	return &SetIsSelectedLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetIsSelectedLotteryLogic) SetIsSelectedLottery(in *pb.SetIsSelectedLotteryReq) (*pb.SetIsSelectedLotteryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SetIsSelectedLotteryResp{}, nil
}
