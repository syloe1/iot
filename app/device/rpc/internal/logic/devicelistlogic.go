package logic

import (
	"context"
	"time"

	"iot-platform/app/device/rpc/internal/model"
	"iot-platform/app/device/rpc/internal/svc"
	"iot-platform/app/device/rpc/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeviceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceListLogic {
	return &DeviceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 设备列表
func (l *DeviceListLogic) DeviceList(in *device.DeviceListReq) (*device.DeviceListResp, error) {
	list, total, err := l.svcCtx.DeviceModel.FindList(l.ctx, in.ProductId, in.Status, in.Keyword, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}

	respList := make([]*device.DeviceInfo, 0, len(list))
	for _, item := range list {
		respList = append(respList, convertDevice(item))
	}

	return &device.DeviceListResp{
		Total: total,
		List:  respList,
	}, nil
}

func convertDevice(item *model.Device) *device.DeviceInfo {
	lastOnlineTime := ""
	if item.LastOnlineTime.Valid {
		lastOnlineTime = item.LastOnlineTime.Time.Format(time.DateTime)
	}

	return &device.DeviceInfo{
		Id:             item.Id,
		ProductId:      item.ProductId,
		Name:           item.Name,
		DeviceKey:      item.DeviceKey,
		Status:         int32(item.Status),
		LastOnlineTime: lastOnlineTime,
		CreatedAt:      item.CreatedAt.Format(time.DateTime),
		UpdatedAt:      item.UpdatedAt.Format(time.DateTime),
	}
}
