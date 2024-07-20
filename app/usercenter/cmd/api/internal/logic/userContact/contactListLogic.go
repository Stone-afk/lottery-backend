package userContact

import (
	"context"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContactListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContactListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContactListLogic {
	return &ContactListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContactListLogic) ContactList(req *types.ContactListReq) (resp *types.ContactListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
