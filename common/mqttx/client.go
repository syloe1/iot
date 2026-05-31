package mqttx

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MustNewClient(c Config) mqtt.Client {
	// 1. 初始化客户端配置项
	opts := mqtt.NewClientOptions().
		AddBroker(c.Broker).                // 指定 MQTT 服务地址
		SetClientID(c.ClientId).            // 设置客户端唯一ID
		SetConnectTimeout(5 * time.Second). // 连接超时 5s
		SetAutoReconnect(true)              // 开启自动重连

	// 2. 按需设置账号密码（非必填）
	if c.Username != "" {
		opts.SetUsername(c.Username)
	}
	if c.Password != "" {
		opts.SetPassword(c.Password)
	}

	// 3. 创建客户端实例
	client := mqtt.NewClient(opts)

	// 4. 发起连接，限时等待 10s
	token := client.Connect()
	if !token.WaitTimeout(10 * time.Second) {
		panic(fmt.Errorf("mqtt connect timeout: %s", c.Broker))
	}

	// 5. 校验连接错误
	if err := token.Error(); err != nil {
		panic(err)
	}

	return client
}
func MustPublish(client mqtt.Client, topic string, qos byte, retained bool, payload []byte) {
	// 发布消息：主题、QoS、是否保留消息、消息体
	token := client.Publish(topic, qos, retained, payload)
	// 阻塞等待发布完成
	token.Wait()
	// 发布失败直接 panic
	if err := token.Error(); err != nil {
		panic(err)
	}
}
