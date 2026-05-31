package svc

import (
	"iot-platform/app/product/rpc/internal/config"
	"iot-platform/app/product/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
