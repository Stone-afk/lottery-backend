package logic

import (
	"context"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveNoticeSubscribePreferenceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveNoticeSubscribePreferenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveNoticeSubscribePreferenceLogic {
	return &SaveNoticeSubscribePreferenceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveNoticeSubscribePreferenceLogic) SaveNoticeSubscribePreference(in *pb.SaveNoticeSubscribePreferenceReq) (*pb.SaveNoticeSubscribePreferenceResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SaveNoticeSubscribePreferenceResp{}, nil
}
