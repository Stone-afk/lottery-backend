package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/xerr"
)

type DelUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserContactLogic {
	return &DelUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserContactLogic) DelUserContact(in *pb.DelUserContactReq) (*pb.DelUserContactResp, error) {
	for _, id := range in.Id {
		err := l.svcCtx.UserContactModel.Delete(l.ctx, id)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db user_contact Delete err:%v, id:%+v", err, in.Id)
		}
	}
	return &pb.DelUserContactResp{}, nil
}
