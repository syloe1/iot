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

type DeviceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceDeleteLogic {
	return &DeviceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceDeleteLogic) DeviceDelete(req *types.DeviceDeleteReq) (resp *types.DeviceDeleteResp, err error) {
	rpcResp, err := l.svcCtx.DeviceRpc.DeviceDelete(l.ctx, &device.DeviceDeleteReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeviceDeleteResp{
		Success: rpcResp.Success,
	}, nil
}
