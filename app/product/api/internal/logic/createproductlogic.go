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

type CreateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProductLogic) CreateProduct(req *types.CreateProductReq) (resp *types.CreateProductResp, err error) {
	rpcResp, err := l.svcCtx.ProductRpc.CreateProduct(l.ctx, &product.CreateProductReq{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateProductResp{
		Success:   rpcResp.Success,
		ProductId: rpcResp.ProductId,
		Msg:       rpcResp.Msg,
	}, nil
}
