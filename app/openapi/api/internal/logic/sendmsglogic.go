// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"errors"

	"iot-platform/app/device/rpc/types/device"
	"iot-platform/app/openapi/api/internal/svc"
	"iot-platform/app/openapi/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMsgLogic) SendMsg(req *types.SendMsgReq) (resp *types.SendMsgResp, err error) {
	if req.DeviceKey == "" {
		return nil, errors.New("device_key is required")
	}
	if req.Payload == "" {
		return nil, errors.New("payload is required")
	}

	_, err = l.svcCtx.DeviceRpc.SendMessageToDevice(l.ctx, &device.SendMessageToDeviceReq{
		DeviceKey: req.DeviceKey,
		Payload:   []byte(req.Payload),
	})
	if err != nil {
		return nil, err
	}

	return &types.SendMsgResp{
		Success: true,
		Msg:     "send message success",
	}, nil
}
