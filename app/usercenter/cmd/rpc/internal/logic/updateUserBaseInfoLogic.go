package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserBaseInfoLogic {
	return &UpdateUserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserBaseInfoLogic) UpdateUserBaseInfo(in *pb.UpdateUserBaseInfoReq) (*pb.UpdateUserBaseInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserBaseInfoResp{}, nil
}
