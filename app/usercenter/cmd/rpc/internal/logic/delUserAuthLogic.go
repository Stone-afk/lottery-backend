package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserAuthLogic {
	return &DelUserAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserAuthLogic) DelUserAuth(in *pb.DelUserAuthReq) (*pb.DelUserAuthResp, error) {
	err := l.svcCtx.UserModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db user Delete err:%v, id:%+v", err, in.Id)
	}
	return &pb.DelUserAuthResp{}, nil
}
