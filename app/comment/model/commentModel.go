package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		CommentList(ctx context.Context, limit, lastId, sort int64) ([]*Comment, error)
		UpdatePraiseNum(ctx context.Context, id, num int64) (int64, error)
		DeleteSoft(ctx context.Context, data *Comment) error
		GetCommentLastId() (int64, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

func (c *customCommentModel) CommentList(ctx context.Context, limit, lastId, sort int64) ([]*Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customCommentModel) UpdatePraiseNum(ctx context.Context, id, num int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customCommentModel) DeleteSoft(ctx context.Context, data *Comment) error {
	//TODO implement me
	panic("implement me")
}

func (c *customCommentModel) GetCommentLastId() (int64, error) {
	//TODO implement me
	panic("implement me")
}

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c, opts...),
	}
}
