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

type SearchUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserSponsorLogic {
	return &SearchUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserSponsorLogic) SearchUserSponsor(in *pb.SearchUserSponsorReq) (*pb.SearchUserSponsorResp, error) {
	list, err := l.svcCtx.UserSponsorModel.FindPageByUserId(l.ctx, in.UserId, in.Page, in.Limit)
	//return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "参数检查 : %v , in :%+v", err, in)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get user sponsor err : %v , in :%+v", err, in)
	}

	var resp []*pb.UserSponsor
	if len(list) > 0 {
		for _, sponsor := range list {
			var pbSponsor pb.UserSponsor
			_ = copier.Copy(&pbSponsor, sponsor)
			resp = append(resp, &pbSponsor)
		}
	}
	return &pb.SearchUserSponsorResp{
		UserSponsor: resp,
	}, nil
}
