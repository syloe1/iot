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

type DeleteProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductLogic) DeleteProduct(req *types.DeleteProductReq) (resp *types.DeleteProductResp, err error) {
	rpcResp, err := l.svcCtx.ProductRpc.DeleteProduct(l.ctx, &product.DeleteProductReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteProductResp{
		Success: rpcResp.Success,
		Msg:     rpcResp.Msg,
	}, nil
}
