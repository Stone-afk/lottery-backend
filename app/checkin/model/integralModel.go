package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IntegralModel = (*customIntegralModel)(nil)

type (
	// IntegralModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIntegralModel.
	IntegralModel interface {
		integralModel
		FindOneByUserId(ctx context.Context, userId int64) (*Integral, error)
		TransUpdateByUserId(ctx context.Context, session sqlx.Session, data *Integral) error
		TransInsertByUserId(ctx context.Context, session sqlx.Session, data *Integral) (sql.Result, error)
		UpdateByUserId(ctx context.Context, data *Integral) error
	}

	customIntegralModel struct {
		*defaultIntegralModel
	}
)

func (m *defaultIntegralModel) FindOneByUserId(ctx context.Context, userId int64) (*Integral, error) {
	//TODO implement me
	panic("implement me")
}

func (m *defaultIntegralModel) TransUpdateByUserId(ctx context.Context, session sqlx.Session, data *Integral) error {
	//TODO implement me
	panic("implement me")
}

func (m *defaultIntegralModel) TransInsertByUserId(ctx context.Context, session sqlx.Session, data *Integral) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (m *defaultIntegralModel) UpdateByUserId(ctx context.Context, data *Integral) error {
	//TODO implement me
	panic("implement me")
}

// NewIntegralModel returns a model for the database table.
func NewIntegralModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) IntegralModel {
	return &customIntegralModel{
		defaultIntegralModel: newIntegralModel(conn, c, opts...),
	}
}
