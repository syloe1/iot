// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"iot-platform/app/push/api/internal/config"
	"iot-platform/common/redisx"

	goredis "github.com/redis/go-redis/v9"
)

type ServiceContext struct {
	Config config.Config
	Redis  *goredis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  redisx.NewClient(c.Redis),
	}
}
