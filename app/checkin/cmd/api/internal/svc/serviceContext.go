package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/checkin/cmd/api/internal/config"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	CheckinRpc    checkin.Checkin
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CheckinRpc:    checkin.NewCheckin(zrpc.MustNewClient(c.CheckinRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}
