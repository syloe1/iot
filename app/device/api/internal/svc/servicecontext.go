// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"iot-platform/app/device/api/internal/config"
	"iot-platform/app/device/rpc/devicesvc"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	DeviceRpc devicesvc.DeviceSvc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		DeviceRpc: devicesvc.NewDeviceSvc(zrpc.MustNewClient(c.DeviceRpc)),
	}
}
