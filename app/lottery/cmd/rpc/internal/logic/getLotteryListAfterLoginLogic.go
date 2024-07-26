package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryListAfterLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryListAfterLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryListAfterLoginLogic {
	return &GetLotteryListAfterLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryListAfterLoginLogic) GetLotteryListAfterLogin(in *pb.GetLotteryListAfterLoginReq) (*pb.GetLotteryListAfterLoginResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetLotteryListAfterLoginResp{}, nil
}
