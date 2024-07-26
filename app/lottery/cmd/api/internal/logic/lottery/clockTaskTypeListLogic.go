package lottery

import (
	"context"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClockTaskTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClockTaskTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClockTaskTypeListLogic {
	return &ClockTaskTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClockTaskTypeListLogic) ClockTaskTypeList(req *types.ClockTaskTypeListReq) (resp *types.ClockTaskTypeListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
