package vote

import (
	"context"

	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVoteRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateVoteRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVoteRecordLogic {
	return &CreateVoteRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateVoteRecordLogic) CreateVoteRecord(req *types.CreateVoteRecordReq) (resp *types.CreateVoteRecordResp, err error) {
	// todo: add your logic here and delete this line

	return
}
