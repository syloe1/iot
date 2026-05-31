package logic

import (
	"context"

	"iot-platform/app/product/rpc/internal/model"
	"iot-platform/app/product/rpc/internal/svc"
	"iot-platform/app/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *product.UpdateProductReq) (*product.UpdateProductResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.ProductModel.Update(l.ctx, &model.Product{
		Id:          in.Id,
		Name:        in.Name,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}
	return &product.UpdateProductResp{
		Success: true,
		Msg:     "update product success",
	}, nil
}
