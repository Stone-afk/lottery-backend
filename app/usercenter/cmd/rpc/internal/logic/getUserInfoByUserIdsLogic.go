package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByUserIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByUserIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByUserIdsLogic {
	return &GetUserInfoByUserIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户表-----------------------
func (l *GetUserInfoByUserIdsLogic) GetUserInfoByUserIds(in *pb.GetUserInfoByUserIdsReq) (*pb.GetUserInfoByUserIdsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserInfoByUserIdsResp{}, nil
}
