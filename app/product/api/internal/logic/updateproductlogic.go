// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"iot-platform/app/product/api/internal/svc"
	"iot-platform/app/product/api/internal/types"
	"iot-platform/app/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductLogic) UpdateProduct(req *types.UpdateProductReq) (resp *types.UpdateProductResp, err error) {
	rpcResp, err := l.svcCtx.ProductRpc.UpdateProduct(l.ctx, &product.UpdateProductReq{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateProductResp{
		Success: rpcResp.Success,
		Msg:     rpcResp.Msg,
	}, nil
}
