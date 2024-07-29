package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PrizeModel = (*customPrizeModel)(nil)

type (
	// PrizeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizeModel.
	PrizeModel interface {
		prizeModel
		//自定义的方法写道这里，避免覆盖
		TransInsert(ctx context.Context, session sqlx.Session, data *Prize) (sql.Result, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		FindByLotteryId(ctx context.Context, lotteryId int64) ([]*Prize, error)
		FindPageByLotteryId(ctx context.Context, lotteryId int64, offset int64, limit int64) ([]*Prize, error)
		GetPrizeInfoByPrizeIds(ctx context.Context, prizeIds []int64) ([]*Prize, error)
		FindAllByLotteryIds(ctx context.Context, lotteryIds []int64) ([]*Prize, error)
	}

	customPrizeModel struct {
		*defaultPrizeModel
	}
)

func (c *customPrizeModel) FindByLotteryId(ctx context.Context, lotteryId int64) ([]*Prize, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customPrizeModel) FindPageByLotteryId(ctx context.Context, lotteryId int64, offset int64, limit int64) ([]*Prize, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customPrizeModel) GetPrizeInfoByPrizeIds(ctx context.Context, prizeIds []int64) ([]*Prize, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customPrizeModel) FindAllByLotteryIds(ctx context.Context, lotteryIds []int64) ([]*Prize, error) {
	//TODO implement me
	panic("implement me")
}

// NewPrizeModel returns a model for the database table.
func NewPrizeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PrizeModel {
	return &customPrizeModel{
		defaultPrizeModel: newPrizeModel(conn, c, opts...),
	}
}
