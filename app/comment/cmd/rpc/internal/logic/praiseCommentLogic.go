package logic

import (
	"context"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PraiseCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPraiseCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PraiseCommentLogic {
	return &PraiseCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PraiseCommentLogic) PraiseComment(in *pb.PraiseCommentReq) (*pb.PraiseCommentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PraiseCommentResp{}, nil
}
