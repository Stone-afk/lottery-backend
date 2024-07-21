package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserDynamicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserDynamicLogic {
	return &SearchUserDynamicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserDynamicLogic) SearchUserDynamic(in *pb.SearchUserDynamicReq) (*pb.SearchUserDynamicResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUserDynamicResp{}, nil
}
