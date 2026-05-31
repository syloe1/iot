package config

import (
	"iot-platform/common/mqttx"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	Name  string
	Mysql struct {
		DataSource string
	}
	Mqtt  mqttx.Config
	Redis redis.RedisConf
}
