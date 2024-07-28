package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ClockTaskRecordModel = (*customClockTaskRecordModel)(nil)

type (
	// ClockTaskRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClockTaskRecordModel.
	ClockTaskRecordModel interface {
		clockTaskRecordModel
		// GetClockTaskRecordByLotteryIdAndUserIds 自定义方法
		// 传入lotteryId以及userIds，得到每个参与者的中奖倍率
		GetClockTaskRecordByLotteryIdAndUserIds(lotteryId int64, userIds []int64) ([]*ClockTaskRecord, error)
	}

	customClockTaskRecordModel struct {
		*defaultClockTaskRecordModel
	}
)

func (c *customClockTaskRecordModel) GetClockTaskRecordByLotteryIdAndUserIds(lotteryId int64, userIds []int64) ([]*ClockTaskRecord, error) {
	//TODO implement me
	panic("implement me")
}

// NewClockTaskRecordModel returns a model for the database table.
func NewClockTaskRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ClockTaskRecordModel {
	return &customClockTaskRecordModel{
		defaultClockTaskRecordModel: newClockTaskRecordModel(conn, c, opts...),
	}
}
