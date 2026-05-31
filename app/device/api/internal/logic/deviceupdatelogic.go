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

type DeviceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceUpdateLogic {
	return &DeviceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceUpdateLogic) DeviceUpdate(req *types.DeviceUpdateReq) (resp *types.DeviceUpdateResp, err error) {
	rpcResp, err := l.svcCtx.DeviceRpc.DeviceUpdate(l.ctx, &device.DeviceUpdateReq{
		Id:             req.Id,
		ProductId:      req.ProductId,
		Name:           req.Name,
		DeviceKey:      req.DeviceKey,
		Status:         req.Status,
		LastOnlineTime: req.LastOnlineTime,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeviceUpdateResp{
		Success: rpcResp.Success,
	}, nil
}
