package comment

import (
	"context"

	"looklook/app/comment/cmd/api/internal/svc"
	"looklook/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentDetailLogic {
	return &GetCommentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentDetailLogic) GetCommentDetail(req *types.CommentDetailReq) (resp *types.CommentDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
