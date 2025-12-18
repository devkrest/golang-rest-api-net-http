package middleware

import (
	"github.com/lakhan-purohit/net-http/internal/pkg/constants"
	"log/slog"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestID, _ := r.Context().Value(constants.RequestIDContextKey).(string)

		slog.Info("incoming_request",
			"request_id", requestID,
			"method", r.Method,
			"path", r.URL.Path,
			"remote", r.RemoteAddr,
		)

		next.ServeHTTP(w, r)

		slog.Debug("request_completed",
			"request_id", requestID,
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
		)
	})
}
