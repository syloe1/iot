package logic

import (
	"context"

	"iot-platform/app/product/rpc/internal/svc"
	"iot-platform/app/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductListLogic) GetProductList(in *product.ProductListReq) (*product.ProductListResp, error) {
	// todo: add your logic here and delete this line
	list, total, err := l.svcCtx.ProductModel.FindList(l.ctx, in.Name, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}
	respList := make([]*product.Product, 0, len(list))
	for _, item := range list {
		respList = append(respList, &product.Product{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
		})
	}
	return &product.ProductListResp{
		Success: true,
		Total:   total,
		List:    respList,
		Msg:     "get product list success",
	}, nil
}
