package config

import (
	"iot-platform/common/mqttx"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	Mqtt mqttx.Config
}
