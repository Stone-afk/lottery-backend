package logic

import (
	"context"

	"looklook/app/upload/cmd/rpc/internal/svc"
	"looklook/app/upload/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadLogic) Upload(in *pb.FileUploadReq) (*pb.FileUploadResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileUploadResp{}, nil
}
