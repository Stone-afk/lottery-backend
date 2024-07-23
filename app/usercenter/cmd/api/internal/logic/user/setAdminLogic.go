package user

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAdminLogic {
	return &SetAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetAdminLogic) SetAdmin(req *types.SetAdminReq) (resp *types.SetAdminResp, err error) {
	_, err = l.svcCtx.UsercenterRpc.SetAdmin(l.ctx, &usercenter.SetAdminReq{
		UserId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}
