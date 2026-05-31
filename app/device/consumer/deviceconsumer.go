package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"iot-platform/app/device/consumer/internal/config"
	"iot-platform/common/mqttx"
	"iot-platform/common/redisx"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const deviceStatusChannel = "device.status"

var configFile = flag.String("f", "etc/device-consumer.yaml", "the config file")

type statusEvent struct {
	DeviceKey string `json:"device_key"`
	Status    int64  `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	db := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redisx.NewClient(c.Redis)
	defer redisClient.Close()

	mqttClient := mqttx.MustNewClient(c.Mqtt)
	defer mqttClient.Disconnect(250)

	subscribe(mqttClient, "device/+/heartbeat", func(client mqtt.Client, msg mqtt.Message) {
		deviceKey := parseDeviceKey(msg.Topic())
		if deviceKey == "" {
			return
		}

		if err := markStatus(context.Background(), db, redisClient, deviceKey, 1); err != nil {
			logx.Errorf("mark device online failed, deviceKey=%s err=%v", deviceKey, err)
		}
	})

	subscribe(mqttClient, "device/+/status", func(client mqtt.Client, msg mqtt.Message) {
		deviceKey := parseDeviceKey(msg.Topic())
		if deviceKey == "" {
			return
		}

		status := int64(0)
		if strings.TrimSpace(string(msg.Payload())) == "online" {
			status = 1
		}

		if err := markStatus(context.Background(), db, redisClient, deviceKey, status); err != nil {
			logx.Errorf("mark device status failed, deviceKey=%s status=%d err=%v", deviceKey, status, err)
		}
	})

	fmt.Printf("Starting device consumer, mqtt=%s...\n", c.Mqtt.Broker)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func subscribe(client mqtt.Client, topic string, handler mqtt.MessageHandler) {
	token := client.Subscribe(topic, 1, handler)
	token.Wait()
	if err := token.Error(); err != nil {
		panic(err)
	}
}

func parseDeviceKey(topic string) string {
	parts := strings.Split(topic, "/")
	if len(parts) < 3 || parts[0] != "device" {
		return ""
	}

	return parts[1]
}

func markStatus(ctx context.Context, db sqlx.SqlConn, redisClient *goredis.Client, deviceKey string, status int64) error {
	lastOnlineTime := sql.NullTime{}
	if status == 1 {
		lastOnlineTime = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	_, err := db.ExecCtx(ctx, "update `device` set `status` = ?, `last_online_time` = ? where `device_key` = ?", status, lastOnlineTime, deviceKey)
	if err != nil {
		return err
	}

	event := statusEvent{
		DeviceKey: deviceKey,
		Status:    status,
		Timestamp: time.Now().Unix(),
	}
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return redisClient.Publish(ctx, deviceStatusChannel, data).Err()
}
