package svc

import (
	"iot-platform/app/user/rpc/internal/config"
	"iot-platform/app/user/rpc/model"
	"iot-platform/common/util"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	JwtAuth   *util.JwtAuth
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
		JwtAuth:   util.NewJwtAuth(c.JwtAuth.AccessSecret, c.JwtAuth.AccessExpire),
	}
}
