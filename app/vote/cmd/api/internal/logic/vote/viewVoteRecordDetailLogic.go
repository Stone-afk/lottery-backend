package vote

import (
	"context"

	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewVoteRecordDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewVoteRecordDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewVoteRecordDetailLogic {
	return &ViewVoteRecordDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewVoteRecordDetailLogic) ViewVoteRecordDetail(req *types.ViewVoteRecordDetailReq) (resp *types.ViewVoteRecordDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
