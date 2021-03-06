// Code generated by goctl. DO NOT EDIT!

package genModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	homestayOrderFieldNames          = builder.RawFieldNames(&HomestayOrder{})
	homestayOrderRows                = strings.Join(homestayOrderFieldNames, ",")
	homestayOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(homestayOrderFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	homestayOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(homestayOrderFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	homestayOrderModel interface {
		Insert(ctx context.Context, data *HomestayOrder) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*HomestayOrder, error)
		Update(ctx context.Context, data *HomestayOrder) error
		Delete(ctx context.Context, id int64) error
	}

	defaultHomestayOrderModel struct {
		conn  sqlx.SqlConn
		table string
	}

	HomestayOrder struct {
		Id int64 `db:"id"`
	}
)

func newHomestayOrderModel(conn sqlx.SqlConn) *defaultHomestayOrderModel {
	return &defaultHomestayOrderModel{
		conn:  conn,
		table: "`homestay_order`",
	}
}

func (m *defaultHomestayOrderModel) Insert(ctx context.Context, data *HomestayOrder) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ()", m.table, homestayOrderRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query)
	return ret, err
}

func (m *defaultHomestayOrderModel) FindOne(ctx context.Context, id int64) (*HomestayOrder, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", homestayOrderRows, m.table)
	var resp HomestayOrder
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHomestayOrderModel) Update(ctx context.Context, data *HomestayOrder) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayOrderRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Id)
	return err
}

func (m *defaultHomestayOrderModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultHomestayOrderModel) tableName() string {
	return m.table
}
