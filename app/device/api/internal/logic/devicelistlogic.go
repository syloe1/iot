// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"iot-platform/app/device/api/internal/svc"
	"iot-platform/app/device/api/internal/types"
	"iot-platform/app/device/rpc/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceListLogic {
	return &DeviceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceListLogic) DeviceList(req *types.DeviceListReq) (resp *types.DeviceListResp, err error) {
	rpcResp, err := l.svcCtx.DeviceRpc.DeviceList(l.ctx, &device.DeviceListReq{
		Page:      req.Page,
		PageSize:  req.PageSize,
		ProductId: req.ProductId,
		Status:    req.Status,
		Keyword:   req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	list := make([]types.DeviceItem, 0, len(rpcResp.List))
	for _, item := range rpcResp.List {
		list = append(list, types.DeviceItem{
			Id:             item.Id,
			ProductId:      item.ProductId,
			Name:           item.Name,
			DeviceKey:      item.DeviceKey,
			Status:         item.Status,
			LastOnlineTime: item.LastOnlineTime,
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
		})
	}

	return &types.DeviceListResp{
		Total: rpcResp.Total,
		List:  list,
	}, nil
}
