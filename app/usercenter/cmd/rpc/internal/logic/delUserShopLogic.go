package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserShopLogic {
	return &DelUserShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserShopLogic) DelUserShop(in *pb.DelUserShopReq) (*pb.DelUserShopResp, error) {
	err := l.svcCtx.UserShopModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db user_shop Delete err:%v, id:%+v", err, in.Id)
	}
	return &pb.DelUserShopResp{}, nil
}
