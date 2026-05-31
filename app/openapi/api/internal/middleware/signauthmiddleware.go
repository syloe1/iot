// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package middleware

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"

	"iot-platform/app/openapi/api/internal/config"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type SignAuthMiddleware struct {
	conf  config.SignAuthConf
	nonce sync.Map
}

func NewSignAuthMiddleware(conf config.SignAuthConf) *SignAuthMiddleware {
	if conf.ExpireSeconds <= 0 {
		conf.ExpireSeconds = 300
	}

	return &SignAuthMiddleware{
		conf: conf,
	}
}

func (m *SignAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := m.verify(r); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		next(w, r)
	}
}

func (m *SignAuthMiddleware) verify(r *http.Request) error {
	appId := r.Header.Get("X-App-Id")
	timestamp := r.Header.Get("X-Timestamp")
	nonce := r.Header.Get("X-Nonce")
	sign := r.Header.Get("X-Sign")
	if appId == "" || timestamp == "" || nonce == "" || sign == "" {
		return fmt.Errorf("missing signature headers")
	}

	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp")
	}

	now := time.Now().Unix()
	if ts < now-m.conf.ExpireSeconds || ts > now+m.conf.ExpireSeconds {
		return fmt.Errorf("timestamp expired")
	}

	secret, ok := m.conf.Apps[appId]
	if !ok || secret == "" {
		return fmt.Errorf("invalid app id")
	}

	nonceKey := appId + ":" + nonce
	if _, loaded := m.nonce.LoadOrStore(nonceKey, now); loaded {
		return fmt.Errorf("nonce reused")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.Body = io.NopCloser(bytes.NewReader(body))

	expected := makeSign(map[string]string{
		"appId":     appId,
		"timestamp": timestamp,
		"nonce":     nonce,
		"body":      string(body),
	}, secret)
	if sign != expected {
		return fmt.Errorf("invalid sign")
	}

	return nil
}

func makeSign(params map[string]string, secret string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for i, key := range keys {
		if i > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(key)
		buf.WriteByte('=')
		buf.WriteString(params[key])
	}
	buf.WriteString("&secret=")
	buf.WriteString(secret)

	sum := sha256.Sum256(buf.Bytes())
	return hex.EncodeToString(sum[:])
}
