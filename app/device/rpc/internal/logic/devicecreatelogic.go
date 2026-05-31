package logic

import (
	"context"
	"database/sql"

	"iot-platform/app/device/rpc/internal/model"
	"iot-platform/app/device/rpc/internal/svc"
	"iot-platform/app/device/rpc/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeviceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceCreateLogic {
	return &DeviceCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建设备
func (l *DeviceCreateLogic) DeviceCreate(in *device.DeviceCreateReq) (*device.DeviceCreateResp, error) {
	ret, err := l.svcCtx.DeviceModel.Insert(l.ctx, &model.Device{
		ProductId:      in.ProductId,
		Name:           in.Name,
		DeviceKey:      in.DeviceKey,
		Status:         int64(in.Status),
		LastOnlineTime: sql.NullTime{},
	})
	if err != nil {
		return nil, err
	}

	id, err := ret.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &device.DeviceCreateResp{
		Id: id,
	}, nil
}
