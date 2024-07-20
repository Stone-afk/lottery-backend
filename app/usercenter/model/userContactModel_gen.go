// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"looklook/common/globalkey"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userContactFieldNames          = builder.RawFieldNames(&UserContact{})
	userContactRows                = strings.Join(userContactFieldNames, ",")
	userContactRowsExpectAutoSet   = strings.Join(stringx.Remove(userContactFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userContactRowsWithPlaceHolder = strings.Join(stringx.Remove(userContactFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLooklookUsercenterUserContactIdPrefix = "cache:looklookUsercenter:userContact:id:"
)

type (
	userContactModel interface {
		Insert(ctx context.Context, data *UserContact) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *UserContact) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserContact, error)
		Update(ctx context.Context, data *UserContact) error
		List(ctx context.Context, page, limit int64) ([]*UserContact, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *UserContact) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*UserContact, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserContact, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserContact, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserContact, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserContact, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultUserContactModel struct {
		sqlc.CachedConn
		table string
	}

	UserContact struct {
		Id         int64     `db:"id"`
		UserId     int64     `db:"user_id"`
		Content    string    `db:"content"` // content
		Remark     string    `db:"remark"`  // remark
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newUserContactModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserContactModel {
	return &defaultUserContactModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_contact`",
	}
}

func (m *defaultUserContactModel) Delete(ctx context.Context, id int64) error {
	looklookUsercenterUserContactIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserContactIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, looklookUsercenterUserContactIdKey)
	return err
}

func (m *defaultUserContactModel) FindOne(ctx context.Context, id int64) (*UserContact, error) {
	looklookUsercenterUserContactIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserContactIdPrefix, id)
	var resp UserContact
	err := m.QueryRowCtx(ctx, &resp, looklookUsercenterUserContactIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userContactRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil,  ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserContactModel) Insert(ctx context.Context, data *UserContact) (sql.Result, error) {
	looklookUsercenterUserContactIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserContactIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userContactRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Content, data.Remark)
	}, looklookUsercenterUserContactIdKey)
	return ret, err
}

func (m *defaultUserContactModel) TransInsert(ctx context.Context, session sqlx.Session, data *UserContact) (sql.Result, error) {
	looklookUsercenterUserContactIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserContactIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userContactRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.Content, data.Remark)
	}, looklookUsercenterUserContactIdKey)
	return ret, err
}
func (m *defaultUserContactModel) Update(ctx context.Context, data *UserContact) error {
	looklookUsercenterUserContactIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserContactIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userContactRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.Content, data.Remark, data.Id)
	}, looklookUsercenterUserContactIdKey)
	return err
}

func (m *defaultUserContactModel) TransUpdate(ctx context.Context, session sqlx.Session, data *UserContact) error {
	looklookUsercenterUserContactIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserContactIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userContactRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.Content, data.Remark, data.Id)
	}, looklookUsercenterUserContactIdKey)
	return err
}

func (m *defaultUserContactModel) List(ctx context.Context, page, limit int64) ([]*UserContact, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", userContactRows, m.table)
	var resp []*UserContact
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultUserContactModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUserContactModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultUserContactModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultUserContactModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*UserContact, error) {

	builder = builder.Columns(userContactRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserContact
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserContactModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserContact, error) {

	builder = builder.Columns(userContactRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserContact
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserContactModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserContact, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(userContactRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*UserContact
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultUserContactModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserContact, error) {

	builder = builder.Columns(userContactRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserContact
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserContactModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserContact, error) {

	builder = builder.Columns(userContactRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserContact
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserContactModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultUserContactModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLooklookUsercenterUserContactIdPrefix, primary)
}

func (m *defaultUserContactModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userContactRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserContactModel) tableName() string {
	return m.table
}
