package logic

import (
	"context"

	"iot-platform/app/device/rpc/internal/svc"
	"iot-platform/app/device/rpc/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeviceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceDeleteLogic {
	return &DeviceDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除设备
func (l *DeviceDeleteLogic) DeviceDelete(in *device.DeviceDeleteReq) (*device.DeviceDeleteResp, error) {
	err := l.svcCtx.DeviceModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &device.DeviceDeleteResp{
		Success: true,
	}, nil
}
