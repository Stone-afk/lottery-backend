package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CheckinRecordModel = (*customCheckinRecordModel)(nil)

type (
	// CheckinRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCheckinRecordModel.
	CheckinRecordModel interface {
		checkinRecordModel
	}

	customCheckinRecordModel struct {
		*defaultCheckinRecordModel
	}
)

// NewCheckinRecordModel returns a model for the database table.
func NewCheckinRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CheckinRecordModel {
	return &customCheckinRecordModel{
		defaultCheckinRecordModel: newCheckinRecordModel(conn, c, opts...),
	}
}
