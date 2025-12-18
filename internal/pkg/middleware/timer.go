package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Timer(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		w.Header().Set("X-Response-Time", duration.String())

		slog.Info("request_duration",
			"path", r.URL.Path,
			"took", time.Since(start),
		)
	})
}
