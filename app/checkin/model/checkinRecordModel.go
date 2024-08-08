package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CheckinRecordModel = (*customCheckinRecordModel)(nil)

type (
	// CheckinRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCheckinRecordModel.
	CheckinRecordModel interface {
		checkinRecordModel
		FindOneByUserId(ctx context.Context, userId int64) (*CheckinRecord, error)
		TransUpdateByUserId(ctx context.Context, session sqlx.Session, data *CheckinRecord) error
		TransInsertByUserId(ctx context.Context, session sqlx.Session, data *CheckinRecord) (sql.Result, error)
	}

	customCheckinRecordModel struct {
		*defaultCheckinRecordModel
	}
)

func (m *defaultCheckinRecordModel) FindOneByUserId(ctx context.Context, userId int64) (*CheckinRecord, error) {
	//TODO implement me
	panic("implement me")
}

func (m *defaultCheckinRecordModel) TransUpdateByUserId(ctx context.Context, session sqlx.Session, data *CheckinRecord) error {
	//TODO implement me
	panic("implement me")
}

func (m *defaultCheckinRecordModel) TransInsertByUserId(ctx context.Context, session sqlx.Session, data *CheckinRecord) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}

// NewCheckinRecordModel returns a model for the database table.
func NewCheckinRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CheckinRecordModel {
	return &customCheckinRecordModel{
		defaultCheckinRecordModel: newCheckinRecordModel(conn, c, opts...),
	}
}
