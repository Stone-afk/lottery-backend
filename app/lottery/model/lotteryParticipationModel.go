package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
	"time"
)

var _ LotteryParticipationModel = (*customLotteryParticipationModel)(nil)

type Lottery3 struct {
	Id              int64     `db:"id"`
	Time            time.Time `db:"time"`
	ParticipationId int64     `db:"participation_id"`
}

type (
	// LotteryParticipationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryParticipationModel.
	LotteryParticipationModel interface {
		lotteryParticipationModel
		GetParticipationUserIdsByLotteryId(ctx context.Context, LotteryId int64) ([]int64, error)
		UpdateWinners(ctx context.Context, LotteryId, UserId, PrizeId int64) error
		GetParticipatorsCountByLotteryId(ctx context.Context, LotteryId int64) (int64, error)
		CheckIsWonByUserIdAndLotteryId(ctx context.Context, LotteryId, UserId int64) (int64, error)
		GetWonListByUserId(ctx context.Context, UserId, Size, LastId int64) ([]*LotteryParticipation, error)
		GetWonListCountByUserId(ctx context.Context, UserId int64) (int64, error)
		CheckIsParticipatedByUserIdAndLotteryId(ctx context.Context, UserId, LotteryId int64) (int64, error)
		GetParticipatedLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error)
		FindAllByUserId(UserId, LastId, Size, IsAnnounced int64) ([]*Lottery2, error)
		FindWonListByUserId(UserId, LastId, Size, IsAnnounced int64) ([]*Lottery3, error)
		GetLastId(ctx context.Context) (int64, error)
	}

	customLotteryParticipationModel struct {
		*defaultLotteryParticipationModel
	}
)

func (c *customLotteryParticipationModel) GetParticipationUserIdsByLotteryId(ctx context.Context, LotteryId int64) ([]int64, error) {
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE lottery_id = ?", c.table)
	var resp []int64
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, LotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GET_PARTICIPATION_USERIDS_BYLOTTERYID_ERROR), "GetParticipationUserIdsByLotteryId,LotteryId:%v, error: %v", LotteryId, err)
	}
	return resp, nil
}

func (c *customLotteryParticipationModel) UpdateWinners(ctx context.Context, LotteryId, UserId, PrizeId int64) error {
	data, err := c.FindOneByLotteryIdUserId(ctx, LotteryId, UserId)
	if err != nil {
		return err
	}
	gozeroLotteryParticipationIdKey := fmt.Sprintf("%s%v", cacheGozeroLotteryParticipationIdPrefix, data.Id)
	gozeroLotteryParticipationLotteryIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheGozeroLotteryParticipationLotteryIdUserIdPrefix, data.LotteryId, data.UserId)
	query := fmt.Sprintf("update %s set is_won = 1, prize_id = ? where `lottery_id` = ? and `user_id` = ?", c.table)
	_, err = c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		res, err := conn.ExecCtx(ctx, query, PrizeId, LotteryId, UserId)
		if err != nil {
			return nil, err
		}
		return res, nil
	}, gozeroLotteryParticipationIdKey, gozeroLotteryParticipationLotteryIdUserIdKey)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.UPDATE_WINNER_ERROR), "UpdateWinners, PrizeId:%v, LotteryId:%v, UserId:%v, error: %v", PrizeId, LotteryId, UserId, err)
	}
	return nil
}

func (c *customLotteryParticipationModel) GetParticipatorsCountByLotteryId(ctx context.Context, LotteryId int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE lottery_id = ?", c.table)
	var resp int64
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, LotteryId)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.GET_PARTICIPATORS_COUNT_BYLOTTERYID_ERROR), "GetParticipatorsCountByLotteryId, LotteryId:%v, error: %v", LotteryId, err)
	}
	return resp, nil
}

func (c *customLotteryParticipationModel) CheckIsWonByUserIdAndLotteryId(ctx context.Context, LotteryId, UserId int64) (int64, error) {
	query := fmt.Sprintf("SELECT is_won FROM %s WHERE lottery_id = ? AND user_id = ?", c.table)
	var resp int64
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, LotteryId, UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.CHECK_ISWON_BYUSERID_ANDLOTTERYID_ERROR), "CheckIsWonByUserIdAndLotteryId, LotteryId:%v, UserId:%v, error: %v", LotteryId, UserId, err)
	}
	return resp, nil
}

func (c *customLotteryParticipationModel) GetWonListByUserId(ctx context.Context, UserId, Size, LastId int64) ([]*LotteryParticipation, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = ? AND is_won = 1 AND id > ? LIMIT ?", c.table)
	var resp []*LotteryParticipation
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, UserId, LastId, Size)
	if err == sqlx.ErrNotFound {
		fmt.Println("sqlx.ErrNotFoundfff", err)
		return resp, nil
	}
	if err != nil {
		fmt.Println("asdfasdf", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GET_WONLIST_BYUSERID_ERROR), "GetWonListByUserId, UserId:%v, Size:%v, error: %v", UserId, Size, err)
	}
	return resp, nil
}

func (c *customLotteryParticipationModel) GetWonListCountByUserId(ctx context.Context, UserId int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id = ? AND is_won = 1", c.table)
	var resp int64
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, UserId)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.GET_WONLISTCOUNT_BYUSERID_ERROR), "GetWonListCountByUserId, UserId:%v, error: %v", UserId, err)
	}
	return resp, nil
}

func (c *customLotteryParticipationModel) CheckIsParticipatedByUserIdAndLotteryId(ctx context.Context, UserId, LotteryId int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id = ? AND lottery_id = ?", c.table)
	var resp int64
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, UserId, LotteryId)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.CHECK_ISPARTICIPATED_BYUSERID_ANDLOTTERYID_ERROR), "CheckIsParticipatedByUserIdAndLotteryId, UserId:%v, LotteryId:%v, error: %v", UserId, LotteryId, err)
	}
	return resp, nil
}

func (c *customLotteryParticipationModel) GetParticipatedLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error) {
	query := fmt.Sprintf("SELECT lottery_id FROM %s WHERE user_id = ?", c.table)
	var resp []int64
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GET_PARTICIPATED_LOTTERYIDS_BYUSERID_ERROR), "GetParticipatedLotteryIdsByUserId, UserId:%v, error: %v", UserId, err)
	}
	return resp, nil
}

// FindAllByUserId 获取当前用户所有参与的抽奖信息
func (c *customLotteryParticipationModel) FindAllByUserId(UserId, LastId, Size, IsAnnounced int64) ([]*Lottery2, error) {
	var query string
	var resp []*Lottery2
	// 获取参与+发起的抽奖信息
	query = fmt.Sprintf(`
SELECT lottery_id as id,create_time as time
FROM (
	SELECT id as lottery_id,create_time FROM lottery 
	WHERE user_id = ?
	AND is_announced = ?
	UNION
	SELECT lottery_id,lottery.create_time FROM lottery_participation 
	LEFT JOIN lottery ON lottery.id = lottery_participation.lottery_id
	WHERE lottery_participation.user_id = ?
	AND lottery.is_announced = ?
)
AS a
WHERE lottery_id < ?
ORDER BY lottery_id
DESC LIMIT ?`)
	err := c.QueryRowsNoCache(&resp, query, UserId, IsAnnounced, UserId, IsAnnounced, LastId, Size)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.FIND_ALLBYUSERID_ERROR), "FindAllByUserId, UserId:%v, LastId:%v, Size:%v, IsAnnounced:%v, error: %v", UserId, LastId, Size, IsAnnounced, err)
	}
	return resp, nil
}

// FindWonListByUserId 获取当前用户所有参与的抽奖信息 (只包括中奖的)
func (c *customLotteryParticipationModel) FindWonListByUserId(UserId, LastId, Size, IsAnnounced int64) ([]*Lottery3, error) {
	query := fmt.Sprintf("SELECT l.id,lp.id as participation_id,lp.update_time as time FROM %s lp LEFT JOIN %s l ON lp.lottery_id = l.id WHERE lp.user_id = ? AND lp.is_won = 1 AND lp.id < ? AND l.is_announced = 1 ORDER BY id DESC LIMIT ?", c.table, "lottery")
	var resp []*Lottery3
	err := c.QueryRowsNoCache(&resp, query, UserId, LastId, Size)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.FIND_WONLIST_BYUSERID_ERROR), "FindWonListByUserId, UserId:%v, LastId:%v, Size:%v, IsAnnounced:%v, error: %v", UserId, LastId, Size, IsAnnounced, err)

	}
	return resp, nil
}

func (c *customLotteryParticipationModel) GetLastId(ctx context.Context) (int64, error) {
	var resp int64
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_GETLASTID_ERROR), "GetLastId, error: %v", err)
	}
	return resp, nil
}

// NewLotteryParticipationModel returns a model for the database table.
func NewLotteryParticipationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryParticipationModel {
	return &customLotteryParticipationModel{
		defaultLotteryParticipationModel: newLotteryParticipationModel(conn, c, opts...),
	}
}
