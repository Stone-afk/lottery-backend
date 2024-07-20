package userSponsor

import (
	"context"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSponsorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSponsorLogic {
	return &AddSponsorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSponsorLogic) AddSponsor(req *types.CreateSponsorReq) (resp *types.CreateSponsorResp, err error) {
	// todo: add your logic here and delete this line

	return
}
