package lottery

import (
	"context"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsWinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckIsWinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsWinLogic {
	return &CheckIsWinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckIsWinLogic) CheckIsWin(req *types.CheckIsWinReq) (resp *types.CheckIsWinResp, err error) {
	// todo: add your logic here and delete this line

	return
}
