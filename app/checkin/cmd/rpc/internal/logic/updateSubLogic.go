package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubLogic {
	return &UpdateSubLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSubLogic) UpdateSub(in *pb.UpdateSubReq) (*pb.UpdateSubResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateSubResp{}, nil
}
