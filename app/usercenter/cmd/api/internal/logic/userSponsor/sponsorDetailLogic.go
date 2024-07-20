package userSponsor

import (
	"context"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSponsorDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorDetailLogic {
	return &SponsorDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SponsorDetailLogic) SponsorDetail(req *types.SponosorDetailReq) (resp *types.SponosorDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
