package userContact

import (
	"context"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDateContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateContactLogic {
	return &UpDateContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpDateContactLogic) UpDateContact(req *types.UpDateContactReq) (resp *types.UpDateContactResp, err error) {
	// todo: add your logic here and delete this line

	return
}
