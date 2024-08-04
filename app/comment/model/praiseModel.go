package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PraiseModel = (*customPraiseModel)(nil)

type (
	// PraiseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPraiseModel.
	PraiseModel interface {
		praiseModel
		PraiseList(ctx context.Context, page, limit, lastId int64) ([]*Praise, error)
		IsPraise(ctx context.Context, commentId, userId int64) (int64, error)
		IsPraiseThisWeek(ctx context.Context, userId int64) (bool, error)
		IsPraiseList(ctx context.Context, commentIds []int64, userId int64) ([]int64, error)
		GetLikeCountByCommentIds(ctx context.Context, commentIds []int64) (map[int64]int64, error)
	}

	customPraiseModel struct {
		*defaultPraiseModel
	}
)

func (c *customPraiseModel) PraiseList(ctx context.Context, page, limit, lastId int64) ([]*Praise, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customPraiseModel) IsPraise(ctx context.Context, commentId, userId int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customPraiseModel) IsPraiseThisWeek(ctx context.Context, userId int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customPraiseModel) IsPraiseList(ctx context.Context, commentIds []int64, userId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customPraiseModel) GetLikeCountByCommentIds(ctx context.Context, commentIds []int64) (map[int64]int64, error) {
	//TODO implement me
	panic("implement me")
}

// NewPraiseModel returns a model for the database table.
func NewPraiseModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PraiseModel {
	return &customPraiseModel{
		defaultPraiseModel: newPraiseModel(conn, c, opts...),
	}
}
