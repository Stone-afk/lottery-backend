package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaskRecordModel = (*customTaskRecordModel)(nil)

type (
	// TaskRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskRecordModel.
	TaskRecordModel interface {
		taskRecordModel
		FindByUserId(ctx context.Context, userId int64, builder squirrel.SelectBuilder, orderBy string) ([]*TaskRecord, error)
		FindByUserIdAndTaskId(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error)
		FindByUserIdAndTaskIdByDay(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error)
		FindByUserIdAndTaskIdByWeek(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error)
	}

	customTaskRecordModel struct {
		*defaultTaskRecordModel
	}
)

func (m *defaultTaskRecordModel) FindByUserId(ctx context.Context, userId int64, builder squirrel.SelectBuilder, orderBy string) ([]*TaskRecord, error) {
	//TODO implement me
	panic("implement me")
}

func (m *defaultTaskRecordModel) FindByUserIdAndTaskId(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error) {
	//TODO implement me
	panic("implement me")
}

func (m *defaultTaskRecordModel) FindByUserIdAndTaskIdByDay(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error) {
	//TODO implement me
	panic("implement me")
}

func (m *defaultTaskRecordModel) FindByUserIdAndTaskIdByWeek(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error) {
	//TODO implement me
	panic("implement me")
}

// NewTaskRecordModel returns a model for the database table.
func NewTaskRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TaskRecordModel {
	return &customTaskRecordModel{
		defaultTaskRecordModel: newTaskRecordModel(conn, c, opts...),
	}
}
