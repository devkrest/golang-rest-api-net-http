package handler

import (
	"net/http"

	"github.com/lakhan-purohit/net-http/internal/pkg/db"
	"github.com/lakhan-purohit/net-http/internal/pkg/response"
	"github.com/lakhan-purohit/net-http/internal/rest-api/repository"
	"github.com/lakhan-purohit/net-http/internal/rest-api/service"
)

func UserHandler() *http.ServeMux {

	mux := http.NewServeMux()

	r := repository.NewUserRepository(db.DB)
	mux.HandleFunc("GET /get-list", service.UserGetListHandler(r))
	mux.HandleFunc("GET /get-full-list", service.UserGetFullListHandler(r))

	// Catch-all for professional 404/405
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.NotFound(response.SendParams{
			W:       w,
			Message: "User endpoint not found or invalid method",
		})
	})

	return mux
}
