package userSponsor

import (
	"context"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateSponsorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDateSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateSponsorLogic {
	return &UpDateSponsorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpDateSponsorLogic) UpDateSponsor(req *types.UpdateSponsorReq) (resp *types.UpdateSponsorResp, err error) {
	// todo: add your logic here and delete this line

	return
}
