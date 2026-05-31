package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		withSession(session sqlx.Session) ProductModel
		FindList(ctx context.Context, name string, page, pageSize int64) ([]*Product, int64, error)
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn),
	}
}

func (m *customProductModel) withSession(session sqlx.Session) ProductModel {
	return NewProductModel(sqlx.NewSqlConnFromSession(session))
}
func (m *customProductModel) FindList(ctx context.Context, name string, page, pageSize int64) ([]*Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	var total int64
	var list []*Product

	if name != "" {
		likeName := "%" + name + "%"
		countQuery := "select count(*) from `product` where `name` like ?"
		if err := m.conn.QueryRowCtx(ctx, &total, countQuery, likeName); err != nil {
			return nil, 0, err
		}

		query := "select " + productRows + " from `product` where `name` like ? order by `id` desc limit ? offset ?"
		if err := m.conn.QueryRowsCtx(ctx, &list, query, likeName, pageSize, offset); err != nil {
			return nil, 0, err
		}

		return list, total, nil
	}

	countQuery := "select count(*) from `product`"
	if err := m.conn.QueryRowCtx(ctx, &total, countQuery); err != nil {
		return nil, 0, err
	}

	query := "select " + productRows + " from `product` order by `id` desc limit ? offset ?"
	if err := m.conn.QueryRowsCtx(ctx, &list, query, pageSize, offset); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
