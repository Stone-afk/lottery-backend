package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIntegralByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIntegralByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIntegralByUserIdLogic {
	return &GetIntegralByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIntegralByUserIdLogic) GetIntegralByUserId(in *pb.GetIntegralByUserIdReq) (*pb.GetIntegralByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetIntegralByUserIdResp{}, nil
}
