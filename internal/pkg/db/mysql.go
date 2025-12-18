package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lakhan-purohit/net-http/internal/pkg/config"
)

var DB *sql.DB

func Connect(cfg config.DBConfig) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("mysql open error:", err)
	}

	// 🔥 Connection pool tuning
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// 🔥 Ping with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("mysql ping error:", err)
	}

	DB = db
	log.Println("✅ MySQL connected")
}

func Close() {
	if DB != nil {
		_ = DB.Close()
		log.Println("🛑 MySQL connection closed")
	}
}
