package logic

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserDynamicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserDynamicLogic {
	return &AddUserDynamicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户动态 UserDynamic-----------------------
func (l *AddUserDynamicLogic) AddUserDynamic(in *pb.AddUserDynamicReq) (*pb.AddUserDynamicResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddUserDynamicResp{}, nil
}
