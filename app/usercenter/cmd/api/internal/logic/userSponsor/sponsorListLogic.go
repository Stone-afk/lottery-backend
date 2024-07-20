package userSponsor

import (
	"context"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSponsorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorListLogic {
	return &SponsorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SponsorListLogic) SponsorList(req *types.SponsorListReq) (resp *types.SponsorListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
