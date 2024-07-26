package lottery

import (
	"context"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChanceTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChanceTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChanceTypeListLogic {
	return &ChanceTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChanceTypeListLogic) ChanceTypeList(req *types.ChanceTypeListReq) (resp *types.ChanceTypeListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
