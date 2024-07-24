package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/app/usercenter/model"
	"looklook/common/xerr"
)

type AddUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserSponsorLogic {
	return &AddUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------抽奖发起人联系方式（抽奖赞助商）-----------------------
func (l *AddUserSponsorLogic) AddUserSponsor(in *pb.AddUserSponsorReq) (*pb.AddUserSponsorResp, error) {
	userSponsor := new(model.UserSponsor)
	err := copier.Copy(userSponsor, in)
	if err != nil {
		//todo 优化错误码
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copier : %+v , err: %v", in, err)
	}
	insert, err := l.svcCtx.UserSponsorModel.Insert(l.ctx, userSponsor)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db user_sponsor Insert err:%v, sponsor:%+v", err, userSponsor)
	}
	id, err := insert.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insertResult.LastInsertId err:%v", err)
	}
	return &pb.AddUserSponsorResp{Id: id}, nil
}
