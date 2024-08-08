package checkin

import (
	"context"

	"looklook/app/checkin/cmd/api/internal/svc"
	"looklook/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCheckinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCheckinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCheckinLogic {
	return &GetCheckinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCheckinLogic) GetCheckin(req *types.GetCheckinReq) (resp *types.GetCheckinResp, err error) {
	// todo: add your logic here and delete this line

	return
}
