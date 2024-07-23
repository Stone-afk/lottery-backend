package userWonDynamicComment

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/xerr"

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
	// 获取用户动态列表
	dynamicList, err := l.svcCtx.UsercenterRpc.SearchUserDynamic(l.ctx, &usercenter.SearchUserDynamicReq{
		UserId: req.UserId,
	})
	if err != nil {
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get SearchUserDynamic"), "Failed to get SearchUserDynamic err : %v ,req:%+v", err, req)
	}
	var userDynamicList []types.DynamicInfo
	if len(dynamicList.UserDynamic) > 0 {
		for _, item := range dynamicList.UserDynamic {
			//logx.Error("item:%v", item)
			var t types.DynamicInfo
			_ = copier.Copy(&t, item)
			//t.UpdateTime = strconv.FormatInt(item.UpdateTime, 10)
			userDynamicList = append(userDynamicList, t)
		}
	}

	// 获取累计奖品数量

	// 获取用户晒单列表

	// 4. 组装返回数据
	return &types.UserWonDynamicCommentResp{
		UserDynamicList: userDynamicList,
	}, nil
}
