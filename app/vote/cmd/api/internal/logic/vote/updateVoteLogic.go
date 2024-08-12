package vote

import (
	"context"

	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateVoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVoteLogic {
	return &UpdateVoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateVoteLogic) UpdateVote(req *types.UpdateVoteReq) (resp *types.UpdateVoteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
