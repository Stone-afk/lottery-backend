package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLotteryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryDetailLogic {
	return &LotteryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LotteryDetailLogic) LotteryDetail(in *pb.LotteryDetailReq) (*pb.LotteryDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LotteryDetailResp{}, nil
}
