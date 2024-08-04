package logic

import (
	"context"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserPraiseLogic {
	return &CheckUserPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------others-----------------------
func (l *CheckUserPraiseLogic) CheckUserPraise(in *pb.CheckUserPraiseReq) (*pb.CheckUserPraiseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckUserPraiseResp{}, nil
}
