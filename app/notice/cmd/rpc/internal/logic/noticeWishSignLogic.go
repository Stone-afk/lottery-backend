package logic

import (
	"context"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NoticeWishSignLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNoticeWishSignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeWishSignLogic {
	return &NoticeWishSignLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NoticeWishSignLogic) NoticeWishSign(in *pb.NoticeWishSignInReq) (*pb.NoticeWishSignInResp, error) {
	// todo: add your logic here and delete this line

	return &pb.NoticeWishSignInResp{}, nil
}
