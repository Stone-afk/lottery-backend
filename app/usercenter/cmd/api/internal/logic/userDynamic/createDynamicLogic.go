package userDynamic

import (
	"context"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDynamicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDynamicLogic {
	return &CreateDynamicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDynamicLogic) CreateDynamic(req *types.CreateDynamicReq) (resp *types.CreateDynamicResp, err error) {
	// todo: add your logic here and delete this line

	return
}
