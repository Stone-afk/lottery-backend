package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NoticeSubscribePreferenceModel = (*customNoticeSubscribePreferenceModel)(nil)

type (
	// NoticeSubscribePreferenceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNoticeSubscribePreferenceModel.
	NoticeSubscribePreferenceModel interface {
		noticeSubscribePreferenceModel
		Upsert(ctx context.Context, data *NoticeSubscribePreference) (sql.Result, error)
	}

	customNoticeSubscribePreferenceModel struct {
		*defaultNoticeSubscribePreferenceModel
	}
)

func (c *customNoticeSubscribePreferenceModel) Upsert(ctx context.Context, data *NoticeSubscribePreference) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}

// NewNoticeSubscribePreferenceModel returns a model for the database table.
func NewNoticeSubscribePreferenceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) NoticeSubscribePreferenceModel {
	return &customNoticeSubscribePreferenceModel{
		defaultNoticeSubscribePreferenceModel: newNoticeSubscribePreferenceModel(conn, c, opts...),
	}
}
