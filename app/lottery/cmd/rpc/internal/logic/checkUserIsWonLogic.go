package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserIsWonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserIsWonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserIsWonLogic {
	return &CheckUserIsWonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserIsWonLogic) CheckUserIsWon(in *pb.CheckUserIsWonReq) (*pb.CheckUserIsWonResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckUserIsWonResp{}, nil
}
