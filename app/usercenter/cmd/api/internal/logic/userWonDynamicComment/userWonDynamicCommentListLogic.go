package userWonDynamicComment

import (
	"context"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserWonDynamicCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserWonDynamicCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserWonDynamicCommentListLogic {
	return &UserWonDynamicCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserWonDynamicCommentListLogic) UserWonDynamicCommentList(req *types.UserWonDynamicCommentReq) (resp *types.UserWonDynamicCommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
