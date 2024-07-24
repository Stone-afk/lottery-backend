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

type AddUserShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserShopLogic {
	return &AddUserShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------userShop-----------------------
func (l *AddUserShopLogic) AddUserShop(in *pb.AddUserShopReq) (*pb.AddUserShopResp, error) {
	userShop := new(model.UserShop)
	err := copier.Copy(userShop, in)
	if err != nil {
		//todo 优化错误码
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copier : %+v , err: %v", in, err)
	}
	insert, err := l.svcCtx.UserShopModel.Insert(l.ctx, userShop)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db user_shop Insert err:%v, shop:%+v", err, userShop)
	}
	_, err = insert.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add Contact db user_shop insertResult.LastInsertId err:%v, shop:%+v", err, userShop)
	}
	return &pb.AddUserShopResp{}, nil
}
