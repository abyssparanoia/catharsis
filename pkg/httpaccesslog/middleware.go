package httpaccesslog

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/abyssparanoia/catharsis/pkg/log"
)

// Middleware ... http access log middleware
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			defer func() {
				log.Debugf(ctx, "http log", zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.String("addr", r.RemoteAddr), zap.String("userAgent", r.UserAgent()))
			}()
			next.ServeHTTP(w, r)
		})
	}
}
