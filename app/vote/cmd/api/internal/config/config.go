package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	rest.RestConf
	VoteRpcConf zrpc.RpcClientConf
}
