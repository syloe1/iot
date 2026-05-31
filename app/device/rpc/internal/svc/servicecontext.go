package svc

import (
	"iot-platform/app/device/rpc/internal/config"
	"iot-platform/app/device/rpc/internal/model"
	"iot-platform/common/mqttx"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	DeviceModel model.DeviceModel
	MqttClient  mqtt.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		DeviceModel: model.NewDeviceModel(sqlx.NewMysql(c.Mysql.DataSource)),
		MqttClient:  mqttx.MustNewClient(c.Mqtt),
	}
}
