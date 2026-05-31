package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DeviceModel = (*customDeviceModel)(nil)

type (
	// DeviceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeviceModel.
	DeviceModel interface {
		deviceModel
		withSession(session sqlx.Session) DeviceModel
		FindList(ctx context.Context, productId int64, status int32, keyword string, page, pageSize int64) ([]*Device, int64, error)
		UpdateStatusByDeviceKey(ctx context.Context, deviceKey string, status int64, lastOnlineTime sql.NullTime) error
	}

	customDeviceModel struct {
		*defaultDeviceModel
	}
)

// NewDeviceModel returns a model for the database table.
func NewDeviceModel(conn sqlx.SqlConn) DeviceModel {
	return &customDeviceModel{
		defaultDeviceModel: newDeviceModel(conn),
	}
}

func (m *customDeviceModel) withSession(session sqlx.Session) DeviceModel {
	return NewDeviceModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customDeviceModel) FindList(ctx context.Context, productId int64, status int32, keyword string, page, pageSize int64) ([]*Device, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	where := "where 1=1"
	args := make([]any, 0)

	if productId > 0 {
		where += " and `product_id` = ?"
		args = append(args, productId)
	}
	if status > 0 {
		where += " and `status` = ?"
		args = append(args, status)
	}
	if keyword != "" {
		where += " and (`name` like ? or `device_key` like ?)"
		likeKeyword := "%" + keyword + "%"
		args = append(args, likeKeyword, likeKeyword)
	}

	var total int64
	countQuery := fmt.Sprintf("select count(*) from %s %s", m.table, where)
	if err := m.conn.QueryRowCtx(ctx, &total, countQuery, args...); err != nil {
		return nil, 0, err
	}

	query := fmt.Sprintf("select %s from %s %s order by `id` desc limit ? offset ?", deviceRows, m.table, where)
	queryArgs := append(args, pageSize, offset)
	var list []*Device
	if err := m.conn.QueryRowsCtx(ctx, &list, query, queryArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customDeviceModel) UpdateStatusByDeviceKey(ctx context.Context, deviceKey string, status int64, lastOnlineTime sql.NullTime) error {
	query := fmt.Sprintf("update %s set `status` = ?, `last_online_time` = ? where `device_key` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, status, lastOnlineTime, deviceKey)
	return err
}
