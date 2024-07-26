package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWonListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWonListLogic {
	return &GetWonListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWonListLogic) GetWonList(in *pb.GetWonListReq) (*pb.GetWonListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetWonListResp{}, nil
}
