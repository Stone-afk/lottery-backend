package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryPrizesListByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryPrizesListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryPrizesListByUserIdLogic {
	return &GetLotteryPrizesListByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryPrizesListByUserIdLogic) GetLotteryPrizesListByUserId(in *pb.GetLotteryPrizesListByUserIdReq) (*pb.GetLotteryPrizesListByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetLotteryPrizesListByUserIdResp{}, nil
}
