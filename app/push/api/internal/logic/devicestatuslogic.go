// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"iot-platform/app/push/api/internal/svc"
)

const deviceStatusChannel = "device.status"

type DeviceStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceStatusLogic {
	return &DeviceStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceStatusLogic) DeviceStatus(w http.ResponseWriter, r *http.Request) error {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	pubsub := l.svcCtx.Redis.Subscribe(r.Context(), deviceStatusChannel)
	defer pubsub.Close()

	ch := pubsub.Channel()
	for {
		select {
		case <-r.Context().Done():
			return nil
		case msg := <-ch:
			if msg == nil {
				return nil
			}
			fmt.Fprintf(w, "event: device.status\n")
			fmt.Fprintf(w, "data: %s\n\n", msg.Payload)
			flusher.Flush()
		}
	}
}
