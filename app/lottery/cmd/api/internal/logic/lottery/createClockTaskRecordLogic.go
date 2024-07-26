package lottery

import (
	"context"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateClockTaskRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateClockTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateClockTaskRecordLogic {
	return &CreateClockTaskRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateClockTaskRecordLogic) CreateClockTaskRecord(req *types.CreateClockTaskRecordReq) (resp *types.CreateClockTaskRecordResp, err error) {
	// todo: add your logic here and delete this line

	return
}
