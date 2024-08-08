package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/checkin/cmd/rpc/internal/config"
	"looklook/app/checkin/model"
	"looklook/app/comment/cmd/rpc/comment"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config              config.Config
	RedisClient         *redis.Redis
	CheckinRecordModel  model.CheckinRecordModel
	IntegralModel       model.IntegralModel
	IntegralRecordModel model.IntegralRecordModel
	TaskRecordModel     model.TaskRecordModel
	TasksModel          model.TasksModel
	TaskProgressModel   model.TaskProgressModel

	UserCenterRpc usercenter.Usercenter
	LotteryRpc    lottery.LotteryZrpcClient
	CommentRpc    comment.CommentZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		CheckinRecordModel:  model.NewCheckinRecordModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		IntegralModel:       model.NewIntegralModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		IntegralRecordModel: model.NewIntegralRecordModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		TaskRecordModel:     model.NewTaskRecordModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		TasksModel:          model.NewTasksModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		TaskProgressModel:   model.NewTaskProgressModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),

		UserCenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
		LotteryRpc:    lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		CommentRpc:    comment.NewCommentZrpcClient(zrpc.MustNewClient(c.CommentRpcConf)),
	}
}
