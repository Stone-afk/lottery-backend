package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSelectedLotteryStatisticLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSelectedLotteryStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSelectedLotteryStatisticLogic {
	return &GetSelectedLotteryStatisticLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSelectedLotteryStatisticLogic) GetSelectedLotteryStatistic(in *pb.GetSelectedLotteryStatisticReq) (*pb.GetSelectedLotteryStatisticResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetSelectedLotteryStatisticResp{}, nil
}
