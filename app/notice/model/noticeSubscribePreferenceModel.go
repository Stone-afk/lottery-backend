package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NoticeSubscribePreferenceModel = (*customNoticeSubscribePreferenceModel)(nil)

type (
	// NoticeSubscribePreferenceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNoticeSubscribePreferenceModel.
	NoticeSubscribePreferenceModel interface {
		noticeSubscribePreferenceModel
	}

	customNoticeSubscribePreferenceModel struct {
		*defaultNoticeSubscribePreferenceModel
	}
)

// NewNoticeSubscribePreferenceModel returns a model for the database table.
func NewNoticeSubscribePreferenceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) NoticeSubscribePreferenceModel {
	return &customNoticeSubscribePreferenceModel{
		defaultNoticeSubscribePreferenceModel: newNoticeSubscribePreferenceModel(conn, c, opts...),
	}
}
