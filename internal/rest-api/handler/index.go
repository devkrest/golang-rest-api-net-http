package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lakhan-purohit/net-http/internal/pkg/middleware"
	"github.com/lakhan-purohit/net-http/internal/pkg/response"
	_ "github.com/lakhan-purohit/net-http/internal/rest-api/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

func GetHandler() *http.ServeMux {

	root := http.NewServeMux()

	// Serve Swagger documentation
	root.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // 👈 FORCE JSON TO PREVENT PARSER ERROR
	))

	// Serve Scalar documentation
	root.HandleFunc("/scalar", func(w http.ResponseWriter, r *http.Request) {
		spec, err := os.ReadFile("./docs/swagger.json")
		if err != nil {
			log.Printf("Error reading swagger.json: %v", err)
			http.Error(w, "Error reading swagger documentation", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
<!doctype html>
<html>
  <head>
    <title>Golang API Documentation</title>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1" />
    <style>
      body {
        margin: 0;
      }
    </style>
  </head>
  <body>
    <script
      id="api-reference"
      data-configuration='{
        "theme": "purple",
        "showModels": true,
        "layout": "sidebar"
      }'></script>
    <script>
      const spec = %s;
      document.getElementById('api-reference').setAttribute('data-spec', JSON.stringify(spec));
    </script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
  </body>
</html>
		`, string(spec))
	})

	root.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only match "/" exactly for the Home Page
		if r.URL.Path != "/" {
			response.NotFound(response.SendParams{
				W:       w,
				Message: "Oops! This endpoint doesn't exist.",
			})
			return
		}

		response.Success(response.SendParams{
			W:    w,
			Data: "Welcome to the Home Page!",
		})
	})

	apiV1 := http.NewServeMux()

	// Public routes
	apiV1.Handle("/public/",
		http.StripPrefix("/public",
			middleware.Public(unAuthenticated()),
		))

	// Private routes
	apiV1.Handle("/private/",
		http.StripPrefix("/private",
			middleware.Private(authenticated()),
		))

	// API v1 Catch-all 404
	apiV1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.NotFound(response.SendParams{
			W:       w,
			Message: "Invalid API endpoint",
		})
	})

	// Mount apiV1 routes
	root.Handle("/api/v1/", http.StripPrefix("/api/v1", middleware.Protected(
		apiV1,
	)))

	return root
}

// Group of authenticated routes
func authenticated() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/user/", http.StripPrefix("/user", UserHandler()))

	// Catch-all 404 for Private
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.NotFound(response.SendParams{
			W:       w,
			Message: "Private route not found",
		})
	})

	return mux
}

// Group of unauthenticated routes
func unAuthenticated() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/auth/", http.StripPrefix("/auth", AuthHandler()))

	// Catch-all 404 for Public
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.NotFound(response.SendParams{
			W:       w,
			Message: "Public route not found",
		})
	})

	return mux
}
