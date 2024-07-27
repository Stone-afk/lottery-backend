package logic

import (
	"context"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NoticeLotteryDrawLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNoticeLotteryDrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeLotteryDrawLogic {
	return &NoticeLotteryDrawLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NoticeLotteryDrawLogic) NoticeLotteryDraw(in *pb.NoticeLotteryDrawReq) (*pb.NoticeLotteryDrawResp, error) {
	// todo: add your logic here and delete this line

	return &pb.NoticeLotteryDrawResp{}, nil
}
