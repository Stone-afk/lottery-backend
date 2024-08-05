package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	DB struct {
		DataSource string
	}
	rest.RestConf
	Cache             cache.CacheConf
	CommentRpcConf    zrpc.RpcClientConf
	UserCenterRpcConf zrpc.RpcClientConf
	LotteryRpcConf    zrpc.RpcClientConf
}
