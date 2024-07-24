package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/app/usercenter/model"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserAuthLogic {
	return &AddUserAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户授权表-----------------------
func (l *AddUserAuthLogic) AddUserAuth(in *pb.AddUserAuthReq) (*pb.AddUserAuthResp, error) {
	userAuth := new(model.UserAuth)
	err := copier.Copy(userAuth, in)
	if err != nil {
		//todo 优化错误码
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copier : %+v , err: %v", in, err)
	}

	insertResult, err := l.svcCtx.UserAuthModel.Insert(l.ctx, userAuth)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add userAuth db user_auth Insert err:%v, auth:%+v", err, userAuth)
	}
	_, err = insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add userAuth db user_auth insertResult.LastInsertId err:%v, auth:%+v", err, userAuth)
	}

	return &pb.AddUserAuthResp{}, nil
}
