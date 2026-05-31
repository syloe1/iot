// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"iot-platform/app/push/api/internal/logic"
	"iot-platform/app/push/api/internal/svc"
)

func deviceStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDeviceStatusLogic(r.Context(), svcCtx)
		if err := l.DeviceStatus(w, r); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
	}
}
