// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"iot-platform/app/device/rpc/devicesvc"
	"iot-platform/app/openapi/api/internal/config"
	"iot-platform/app/openapi/api/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	DeviceRpc devicesvc.DeviceSvc
	SignAuth  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		DeviceRpc: devicesvc.NewDeviceSvc(zrpc.MustNewClient(c.DeviceRpc)),
		SignAuth:  middleware.NewSignAuthMiddleware(c.SignAuth).Handle,
	}
}
