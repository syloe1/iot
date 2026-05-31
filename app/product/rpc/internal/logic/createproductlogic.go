package logic

import (
	"context"

	"iot-platform/app/product/rpc/internal/model"
	"iot-platform/app/product/rpc/internal/svc"
	"iot-platform/app/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *product.CreateProductReq) (*product.CreateProductResp, error) {
	// todo: add your logic here and delete this line
	ret, err := l.svcCtx.ProductModel.Insert(l.ctx, &model.Product{
		Name:        in.Name,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &product.CreateProductResp{
		Success:   true,
		ProductId: id,
		Msg:       "create product success",
	}, nil
}
