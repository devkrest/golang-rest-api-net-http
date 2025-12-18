package middleware

import (
	"context"
	"net/http"

	"github.com/lakhan-purohit/net-http/internal/pkg/constants"
	"github.com/lakhan-purohit/net-http/internal/pkg/utils"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = utils.UUID()
		}

		w.Header().Set("X-Request-ID", id)
		ctx := context.WithValue(r.Context(), constants.RequestIDContextKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
