package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryParticipationModel) UpdateWinners(ctx context.Context, LotteryId, UserId, PrizeId int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryParticipationModel) GetParticipatorsCountByLotteryId(ctx context.Context, LotteryId int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryParticipationModel) CheckIsWonByUserIdAndLotteryId(ctx context.Context, LotteryId, UserId int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryParticipationModel) GetWonListByUserId(ctx context.Context, UserId, Size, LastId int64) ([]*LotteryParticipation, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryParticipationModel) GetWonListCountByUserId(ctx context.Context, UserId int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryParticipationModel) CheckIsParticipatedByUserIdAndLotteryId(ctx context.Context, UserId, LotteryId int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryParticipationModel) GetParticipatedLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryParticipationModel) FindAllByUserId(UserId, LastId, Size, IsAnnounced int64) ([]*Lottery2, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryParticipationModel) FindWonListByUserId(UserId, LastId, Size, IsAnnounced int64) ([]*Lottery3, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customLotteryParticipationModel) GetLastId(ctx context.Context) (int64, error) {
	//TODO implement me
	panic("implement me")
}

// NewLotteryParticipationModel returns a model for the database table.
func NewLotteryParticipationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryParticipationModel {
	return &customLotteryParticipationModel{
		defaultLotteryParticipationModel: newLotteryParticipationModel(conn, c, opts...),
	}
}
