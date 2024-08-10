package checkin

import (
	"context"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/common/ctxdata"

	"looklook/app/checkin/cmd/api/internal/svc"
	"looklook/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckinLogic {
	return &CheckinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckinLogic) Checkin(req *types.CheckinReq) (resp *types.CheckinResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	record, err := l.svcCtx.CheckinRpc.UpdateCheckinRecord(l.ctx, &checkin.UpdateCheckinRecordReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.CheckinResp{
		State:                 record.State,
		ContinuousCheckinDays: record.ContinuousCheckinDays,
		Integral:              record.Integral,
	}, nil
}
