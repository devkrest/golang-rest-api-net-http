package handler

import (
	"net/http"

	"github.com/lakhan-purohit/net-http/internal/pkg/db"
	"github.com/lakhan-purohit/net-http/internal/pkg/response"
	"github.com/lakhan-purohit/net-http/internal/rest-api/repository"
	"github.com/lakhan-purohit/net-http/internal/rest-api/service"
)

func AuthHandler() *http.ServeMux {
	mux := http.NewServeMux()

	authRepo := repository.NewAuthRepository(db.DB)
	mux.HandleFunc("POST /login", service.LoginHandler(authRepo))
	mux.HandleFunc("POST /sign-up", service.SignUpHandler(authRepo))

	// Catch-all for professional 404/405
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.NotFound(response.SendParams{
			W:       w,
			Message: "Auth endpoint not found or invalid method",
		})
	})

	return mux
}
