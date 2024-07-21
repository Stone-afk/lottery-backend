package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
)

type AddUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserSponsorLogic {
	return &AddUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------抽奖发起人联系方式（抽奖赞助商）-----------------------
func (l *AddUserSponsorLogic) AddUserSponsor(in *pb.AddUserSponsorReq) (*pb.AddUserSponsorResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddUserSponsorResp{}, nil
}
