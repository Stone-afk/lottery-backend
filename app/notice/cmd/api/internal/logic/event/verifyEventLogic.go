package event

import (
	"context"

	"looklook/app/notice/cmd/api/internal/svc"
	"looklook/app/notice/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyEventLogic {
	return &VerifyEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyEventLogic) VerifyEvent(req *types.VerifyEventReq) (resp *types.VerifyEventResp, err error) {
	// todo: add your logic here and delete this line

	return
}
