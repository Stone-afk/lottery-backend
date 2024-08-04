package comment

import (
	"context"

	"looklook/app/comment/cmd/api/internal/svc"
	"looklook/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentPraiseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPraiseLogic {
	return &CommentPraiseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentPraiseLogic) CommentPraise(req *types.CommentPraiseReq) (resp *types.CommentPraiseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
