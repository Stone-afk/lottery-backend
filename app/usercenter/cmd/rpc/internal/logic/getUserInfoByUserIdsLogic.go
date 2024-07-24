package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/app/usercenter/model"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByUserIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByUserIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByUserIdsLogic {
	return &GetUserInfoByUserIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户表-----------------------
func (l *GetUserInfoByUserIdsLogic) GetUserInfoByUserIds(in *pb.GetUserInfoByUserIdsReq) (*pb.GetUserInfoByUserIdsResp, error) {
	list, err := l.svcCtx.UserModel.FindUserInfoByUserIds(l.ctx, in.UserIds)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get user info  err : %v , in :%+v", err, in)
	}

	var resp []*pb.UserInfoForComment
	if len(list) > 0 {
		for _, userInfo := range list {
			var pbUserInfo pb.UserInfoForComment
			_ = copier.Copy(&pbUserInfo, userInfo)
			resp = append(resp, &pbUserInfo)
		}
	}
	return &pb.GetUserInfoByUserIdsResp{
		UserInfo: resp,
	}, nil
}
