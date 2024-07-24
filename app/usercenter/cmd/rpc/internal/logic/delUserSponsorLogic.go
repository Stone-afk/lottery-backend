package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserSponsorLogic {
	return &DelUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserSponsorLogic) DelUserSponsor(in *pb.DelUserSponsorReq) (*pb.DelUserSponsorResp, error) {
	err := l.svcCtx.UserSponsorModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db user_sponsor Delete err:%v, id:%+v", err, in.Id)
	}
	return &pb.DelUserSponsorResp{}, nil
}
