package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserDynamicModel = (*customUserDynamicModel)(nil)

type (
	// UserDynamicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserDynamicModel.
	UserDynamicModel interface {
		userDynamicModel
	}

	customUserDynamicModel struct {
		*defaultUserDynamicModel
	}
)

// NewUserDynamicModel returns a model for the database table.
func NewUserDynamicModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserDynamicModel {
	return &customUserDynamicModel{
		defaultUserDynamicModel: newUserDynamicModel(conn, c, opts...),
	}
}
