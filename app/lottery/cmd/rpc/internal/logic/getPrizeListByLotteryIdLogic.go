package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeListByLotteryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeListByLotteryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeListByLotteryIdLogic {
	return &GetPrizeListByLotteryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizeListByLotteryIdLogic) GetPrizeListByLotteryId(in *pb.GetPrizeListByLotteryIdReq) (*pb.GetPrizeListByLotteryIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPrizeListByLotteryIdResp{}, nil
}
