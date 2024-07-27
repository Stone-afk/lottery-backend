package logic

import (
	"context"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NoticeLotteryStartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNoticeLotteryStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeLotteryStartLogic {
	return &NoticeLotteryStartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NoticeLotteryStartLogic) NoticeLotteryStart(in *pb.NoticeLotteryStartReq) (*pb.NoticeLotteryStartResp, error) {
	// todo: add your logic here and delete this line

	return &pb.NoticeLotteryStartResp{}, nil
}
