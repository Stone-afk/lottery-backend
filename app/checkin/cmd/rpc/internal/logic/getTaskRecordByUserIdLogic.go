package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskRecordByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskRecordByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskRecordByUserIdLogic {
	return &GetTaskRecordByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskRecordByUserIdLogic) GetTaskRecordByUserId(in *pb.GetTaskRecordByUserIdReq) (*pb.GetTaskRecordByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetTaskRecordByUserIdResp{}, nil
}
