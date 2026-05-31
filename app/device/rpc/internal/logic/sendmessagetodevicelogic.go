package logic

import (
	"context"
	"errors"
	"fmt"

	"iot-platform/app/device/rpc/internal/svc"
	"iot-platform/app/device/rpc/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageToDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageToDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageToDeviceLogic {
	return &SendMessageToDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageToDeviceLogic) SendMessageToDevice(in *device.SendMessageToDeviceReq) (*device.Empty, error) {
	if in.DeviceKey == "" {
		return nil, errors.New("device_key is required")
	}
	if len(in.Payload) == 0 {
		return nil, errors.New("payload is required")
	}

	topic := fmt.Sprintf("device/%s/command", in.DeviceKey)
	token := l.svcCtx.MqttClient.Publish(topic, 1, false, in.Payload)
	token.Wait()
	if err := token.Error(); err != nil {
		return nil, err
	}

	l.Infof("published message to mqtt, topic=%s payloadSize=%d", topic, len(in.Payload))

	return &device.Empty{}, nil
}
