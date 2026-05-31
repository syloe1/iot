package logic

import (
	"context"

	"iot-platform/app/product/rpc/internal/svc"
	"iot-platform/app/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *product.DeleteProductReq) (*product.DeleteProductResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.ProductModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &product.DeleteProductResp{
		Success: true,
		Msg:     "delete product success",
	}, nil
}
