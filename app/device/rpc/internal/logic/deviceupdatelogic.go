package logic

import (
	"context"
	"database/sql"
	"time"

	"iot-platform/app/device/rpc/internal/model"
	"iot-platform/app/device/rpc/internal/svc"
	"iot-platform/app/device/rpc/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeviceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceUpdateLogic {
	return &DeviceUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改设备
func (l *DeviceUpdateLogic) DeviceUpdate(in *device.DeviceUpdateReq) (*device.DeviceUpdateResp, error) {
	lastOnlineTime := sql.NullTime{}
	if in.LastOnlineTime != "" {
		parsed, err := time.ParseInLocation(time.DateTime, in.LastOnlineTime, time.Local)
		if err != nil {
			return nil, err
		}
		lastOnlineTime = sql.NullTime{
			Time:  parsed,
			Valid: true,
		}
	}

	err := l.svcCtx.DeviceModel.Update(l.ctx, &model.Device{
		Id:             in.Id,
		ProductId:      in.ProductId,
		Name:           in.Name,
		DeviceKey:      in.DeviceKey,
		Status:         int64(in.Status),
		LastOnlineTime: lastOnlineTime,
	})
	if err != nil {
		return nil, err
	}

	return &device.DeviceUpdateResp{
		Success: true,
	}, nil
}
