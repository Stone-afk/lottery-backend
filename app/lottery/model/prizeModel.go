package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
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
	var resp []*Prize
	query := fmt.Sprintf("SELECT * FROM %s WHERE lottery_id = ?", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_PRIZES_BYLOTTERYID_ERROR), "QueryRowsNoCacheCtx, &resp:%v, query:%v, lotteryId:%v, error: %v", &resp, query, lotteryId, err)
	}
	return resp, nil
}

func (c *customPrizeModel) FindPageByLotteryId(ctx context.Context, lotteryId int64, offset int64, limit int64) ([]*Prize, error) {
	var resp []*Prize
	query := fmt.Sprintf("SELECT * FROM %s WHERE lottery_id = ? limit ?,?", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lotteryId, (offset-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *customPrizeModel) GetPrizeInfoByPrizeIds(ctx context.Context, prizeIds []int64) ([]*Prize, error) {
	var resp []*Prize
	// 这里传int64类型的切片，需要将切片转换成字符串，然后在sql语句中使用in关键字
	prizeIdsStr := ""
	for i, v := range prizeIds {
		if i == 0 {
			prizeIdsStr = fmt.Sprintf("%d", v)
		} else {
			prizeIdsStr = fmt.Sprintf("%s,%d", prizeIdsStr, v)
		}
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id in (%s)", c.table, prizeIdsStr)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *customPrizeModel) FindAllByLotteryIds(ctx context.Context, lotteryIds []int64) ([]*Prize, error) {
	var resp []*Prize
	// 这里传int64类型的切片，需要将切片转换成字符串，然后在sql语句中使用in关键字
	lotteryIdsStr := ""
	for i, v := range lotteryIds {
		if i == 0 {
			lotteryIdsStr = fmt.Sprintf("%d", v)
		} else {
			lotteryIdsStr = fmt.Sprintf("%s,%d", lotteryIdsStr, v)
		}
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE lottery_id in (%s)", c.table, lotteryIdsStr)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewPrizeModel returns a model for the database table.
func NewPrizeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PrizeModel {
	return &customPrizeModel{
		defaultPrizeModel: newPrizeModel(conn, c, opts...),
	}
}
