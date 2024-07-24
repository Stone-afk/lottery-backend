package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserAddressLogic {
	return &DelUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserAddressLogic) DelUserAddress(in *pb.DelUserAddressReq) (*pb.DelUserAddressResp, error) {
	err := l.svcCtx.UserAddressModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db user_address Delete err:%v, id:%+v", err, in.Id)
	}
	return &pb.DelUserAddressResp{}, nil
}
