package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/lakhan-purohit/net-http/internal/pkg/constants"
	"github.com/lakhan-purohit/net-http/internal/pkg/response"
	"github.com/lakhan-purohit/net-http/internal/pkg/utils"
)

func JWT(next http.Handler) http.Handler {
	jwtService := utils.NewJWT()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			response.UnauthorizedAccess(response.SendParams{
				W:       w,
				Message: "missing token",
			})
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")

		claims, err := jwtService.Parse(token)
		if err != nil {
			response.UnauthorizedAccess(response.SendParams{
				W:       w,
				Message: "invalid token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), constants.UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
