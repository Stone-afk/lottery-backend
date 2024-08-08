package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaskProgressModel = (*customTaskProgressModel)(nil)

type (
	// TaskProgressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskProgressModel.
	TaskProgressModel interface {
		taskProgressModel
		FindOneByUserId(ctx context.Context, userId int64) (*TaskProgress, error)
		InsertByUserId(ctx context.Context, data *TaskProgress) (sql.Result, error)
		TransUpdateByUserId(ctx context.Context, session sqlx.Session, data *TaskProgress) error
		UpdateByUserId(ctx context.Context, data *TaskProgress) error
		FindAllSubId(ctx context.Context) ([]int64, error)
		TransInsertByUserId(ctx context.Context, session sqlx.Session, data *TaskProgress) (sql.Result, error)
	}

	customTaskProgressModel struct {
		*defaultTaskProgressModel
	}
)

func (m *defaultTaskProgressModel) FindOneByUserId(ctx context.Context, userId int64) (*TaskProgress, error) {
	//TODO implement me
	panic("implement me")
}

func (m *defaultTaskProgressModel) InsertByUserId(ctx context.Context, data *TaskProgress) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (m *defaultTaskProgressModel) TransUpdateByUserId(ctx context.Context, session sqlx.Session, data *TaskProgress) error {
	//TODO implement me
	panic("implement me")
}

func (m *defaultTaskProgressModel) UpdateByUserId(ctx context.Context, data *TaskProgress) error {
	//TODO implement me
	panic("implement me")
}

func (m *defaultTaskProgressModel) FindAllSubId(ctx context.Context) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *defaultTaskProgressModel) TransInsertByUserId(ctx context.Context, session sqlx.Session, data *TaskProgress) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}

// NewTaskProgressModel returns a model for the database table.
func NewTaskProgressModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TaskProgressModel {
	return &customTaskProgressModel{
		defaultTaskProgressModel: newTaskProgressModel(conn, c, opts...),
	}
}
