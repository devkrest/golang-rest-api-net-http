package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lakhan-purohit/net-http/internal/pkg/config"
	"github.com/lakhan-purohit/net-http/internal/pkg/db"
	"github.com/lakhan-purohit/net-http/internal/rest-api/handler"
)

func Run() {
	cfg := config.Get()

	srv := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      handler.GetHandler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		log.Printf("🚀 Server started on port %s", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// Graceful shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("🛑 Shutting down server...")

	// Close DB
	db.Close()
	log.Println("🗄️ Database connection closed")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("✅ Server exited gracefully")
}
