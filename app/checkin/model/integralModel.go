package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IntegralModel = (*customIntegralModel)(nil)

type (
	// IntegralModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIntegralModel.
	IntegralModel interface {
		integralModel
	}

	customIntegralModel struct {
		*defaultIntegralModel
	}
)

// NewIntegralModel returns a model for the database table.
func NewIntegralModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) IntegralModel {
	return &customIntegralModel{
		defaultIntegralModel: newIntegralModel(conn, c, opts...),
	}
}
