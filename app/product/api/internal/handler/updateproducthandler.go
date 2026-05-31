// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"iot-platform/app/product/api/internal/logic"
	"iot-platform/app/product/api/internal/svc"
	"iot-platform/app/product/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func updateProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateProductLogic(r.Context(), svcCtx)
		resp, err := l.UpdateProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
