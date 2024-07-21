package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
)

type DelUserDynamicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserDynamicLogic {
	return &DelUserDynamicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserDynamicLogic) DelUserDynamic(in *pb.DelUserDynamicReq) (*pb.DelUserDynamicResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelUserDynamicResp{}, nil
}
