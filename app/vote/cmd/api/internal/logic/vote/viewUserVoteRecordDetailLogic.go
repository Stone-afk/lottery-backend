package vote

import (
	"context"

	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewUserVoteRecordDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewUserVoteRecordDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewUserVoteRecordDetailLogic {
	return &ViewUserVoteRecordDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewUserVoteRecordDetailLogic) ViewUserVoteRecordDetail(req *types.ViewUserVoteRecordDetailReq) (resp *types.ViewUserVoteRecordDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
