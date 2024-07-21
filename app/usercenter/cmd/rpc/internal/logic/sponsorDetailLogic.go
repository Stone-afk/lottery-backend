package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSponsorDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorDetailLogic {
	return &SponsorDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SponsorDetailLogic) SponsorDetail(in *pb.SponsorDetailReq) (*pb.SponsorDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SponsorDetailResp{}, nil
}
