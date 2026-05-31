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

type GetProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductListLogic) GetProductList(req *types.ProductListReq) (resp *types.ProductListResp, err error) {
	rpcResp, err := l.svcCtx.ProductRpc.GetProductList(l.ctx, &product.ProductListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}

	list := make([]types.Product, 0, len(rpcResp.List))
	for _, item := range rpcResp.List {
		list = append(list, types.Product{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
		})
	}

	return &types.ProductListResp{
		Success: rpcResp.Success,
		Total:   rpcResp.Total,
		List:    list,
		Msg:     rpcResp.Msg,
	}, nil
}
