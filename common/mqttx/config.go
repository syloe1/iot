package mqttx

// Config MQTT 客户端连接配置
type Config struct {
	Broker   string // MQTT 服务端地址，格式：ip:port，例：127.0.0.1:1883
	ClientId string // MQTT 客户端唯一标识，每个连接 Broker 的客户端不能重复
	Username string `json:",optional"` // 连接账号，可选字段
	Password string `json:",optional"` // 连接密码，可选字段
}
