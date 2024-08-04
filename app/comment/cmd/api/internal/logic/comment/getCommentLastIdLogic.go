package comment

import (
	"context"

	"looklook/app/comment/cmd/api/internal/svc"
	"looklook/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLastIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentLastIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLastIdLogic {
	return &GetCommentLastIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentLastIdLogic) GetCommentLastId(req *types.GetCommentLastIdReq) (resp *types.GetCommentLastIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
