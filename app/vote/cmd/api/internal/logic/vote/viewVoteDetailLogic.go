package vote

import (
	"context"

	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewVoteDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewVoteDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewVoteDetailLogic {
	return &ViewVoteDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewVoteDetailLogic) ViewVoteDetail(req *types.ViewVoteDetailReq) (resp *types.ViewVoteDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
