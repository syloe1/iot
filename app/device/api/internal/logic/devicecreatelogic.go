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

type DeviceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceCreateLogic {
	return &DeviceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceCreateLogic) DeviceCreate(req *types.DeviceCreateReq) (resp *types.DeviceCreateResp, err error) {
	rpcResp, err := l.svcCtx.DeviceRpc.DeviceCreate(l.ctx, &device.DeviceCreateReq{
		ProductId: req.ProductId,
		Name:      req.Name,
		DeviceKey: req.DeviceKey,
		Status:    req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeviceCreateResp{
		Id: rpcResp.Id,
	}, nil
}
