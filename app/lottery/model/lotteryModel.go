package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ LotteryModel = (*customLotteryModel)(nil)

type Lottery2 struct {
	Id   int64     `db:"id"`
	Time time.Time `db:"time"`
}

type (
	// LotteryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryModel.
	LotteryModel interface {
		lotteryModel
		// 自定义方法
		TransUpdateClockTaskId(ctx context.Context, session sqlx.Session, data *Lottery) (sql.Result, error)
		UpdatePublishTime(ctx context.Context, data *Lottery) error
		LotteryList(ctx context.Context, page, selected, lastId int64) ([]*Lottery, error)
		FindUserIdByLotteryId(ctx context.Context, lotteryId int64) (*int64, error)
		GetLotterysByLessThanCurrentTime(ctx context.Context, currentTime time.Time, announceType int64) ([]int64, error)
		UpdateLotteryStatus(ctx context.Context, lotteryID int64) error
		GetTypeIs2AndIsNotAnnounceLotterys(ctx context.Context, announceType int64) ([]*Lottery, error)
		GetLotteryIdByUserId(ctx context.Context, UserId int64) (*int64, error)
		GetTodayLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error)
		GetWeekLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error)
		// GetLotteryListAfterLogin 登录后获取抽奖列表
		GetLotteryListAfterLogin(ctx context.Context, size, isSelected, lastId int64, lotteryIds []int64) ([]*Lottery, error)
		GetLastId(ctx context.Context) (int64, error)
		FindAllByUserId(UserId, LastId, Size, IsAnnounced int64) ([]*Lottery2, error)
	}

	customLotteryModel struct {
		*defaultLotteryModel
	}
)

func (c *customLotteryModel) TransUpdateClockTaskId(ctx context.Context, session sqlx.Session, data *Lottery) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) UpdatePublishTime(ctx context.Context, data *Lottery) error {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) LotteryList(ctx context.Context, page, selected, lastId int64) ([]*Lottery, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) FindUserIdByLotteryId(ctx context.Context, lotteryId int64) (*int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) GetLotterysByLessThanCurrentTime(ctx context.Context, currentTime time.Time, announceType int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) UpdateLotteryStatus(ctx context.Context, lotteryID int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) GetTypeIs2AndIsNotAnnounceLotterys(ctx context.Context, announceType int64) ([]*Lottery, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) GetLotteryIdByUserId(ctx context.Context, UserId int64) (*int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) GetTodayLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) GetWeekLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) GetLotteryListAfterLogin(ctx context.Context, size, isSelected, lastId int64, lotteryIds []int64) ([]*Lottery, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) GetLastId(ctx context.Context) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryModel) FindAllByUserId(UserId, LastId, Size, IsAnnounced int64) ([]*Lottery2, error) {
	//TODO implement me
	panic("implement me")
}

// NewLotteryModel returns a model for the database table.
func NewLotteryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryModel {
	return &customLotteryModel{
		defaultLotteryModel: newLotteryModel(conn, c, opts...),
	}
}
