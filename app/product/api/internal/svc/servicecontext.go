// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"iot-platform/app/product/api/internal/config"
	"iot-platform/app/product/rpc/productsvc"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc productsvc.ProductSvc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productsvc.NewProductSvc(zrpc.MustNewClient(c.ProductRpc)),
	}
}
