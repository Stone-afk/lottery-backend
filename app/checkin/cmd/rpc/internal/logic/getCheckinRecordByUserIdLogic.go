package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCheckinRecordByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCheckinRecordByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCheckinRecordByUserIdLogic {
	return &GetCheckinRecordByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCheckinRecordByUserIdLogic) GetCheckinRecordByUserId(in *pb.GetCheckinRecordByUserIdReq) (*pb.GetCheckinRecordByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCheckinRecordByUserIdResp{}, nil
}
