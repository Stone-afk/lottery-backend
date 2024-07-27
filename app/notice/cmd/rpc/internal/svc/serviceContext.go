package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/notice/cmd/rpc/internal/config"
	"looklook/app/notice/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config                         config.Config
	AsynqClient                    *asynq.Client
	UserCenterRpc                  usercenter.Usercenter
	LotteryRpc                     lottery.LotteryZrpcClient
	NoticeSubscribePreferenceModel model.NoticeSubscribePreferenceModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                         c,
		AsynqClient:                    newAsynqClient(c),
		UserCenterRpc:                  usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
		LotteryRpc:                     lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		NoticeSubscribePreferenceModel: model.NewNoticeSubscribePreferenceModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
