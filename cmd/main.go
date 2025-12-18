// @title           Golang API Documentation
// @version         1.0
// @description     A robust Authentication & User management API built with Go's net/http.
// @description     ## Features
// @description     - ⚡ **High Performance** - Built with Go net/http
// @description     - 🔐 **JWT Authentication** - Secure user authentication
// @description     - 📁 **File Upload** - Media management with validation
// @description     - 🗄️ **MySQL Database** - Robust data persistence
// @description     - 🚀 **Auto Documentation** - Interactive Scalar & Swagger UI
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8001
// @BasePath        /
// @schemes         http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

package main

import (
	"log/slog"
	"os"

	"github.com/lakhan-purohit/net-http/internal/pkg/config"
	"github.com/lakhan-purohit/net-http/internal/pkg/db"
	"github.com/lakhan-purohit/net-http/internal/pkg/server"
)

func main() {

	config.Load()
	cfg := config.Get()

	// 🔥 Initialize Structured Logging
	var handler slog.Handler
	if cfg.App.Env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	}
	slog.SetDefault(slog.New(handler))

	slog.Info("Starting Golang API", "env", cfg.App.Env, "port", cfg.App.Port)

	// 🔥 Connect DB
	db.Connect(cfg.DB)

	// 🔥 Graceful shutdown
	server.Run()

}
