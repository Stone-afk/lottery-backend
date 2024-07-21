package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxMiniRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWxMiniRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniRegisterLogic {
	return &WxMiniRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WxMiniRegisterLogic) WxMiniRegister(in *pb.WXMiniRegisterReq) (*pb.WXMiniRegisterResp, error) {
	// todo: add your logic here and delete this line

	return &pb.WXMiniRegisterResp{}, nil
}
