// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"iot-platform/app/user/api/internal/config"
	"iot-platform/app/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
