package config

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
	DB  DBConfig
	JWT JWTConfig
}

type AppConfig struct {
	Env  string
	Port string
}

type DBConfig struct {
	Host            string
	Port            string
	Name            string
	User            string
	Password        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type JWTConfig struct {
	Secret            string
	AccessExpiration  time.Duration
	RefreshExpiration time.Duration
}

var cfg *Config

func Load() {
	// Load .env only in dev

	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}

	cfg = &Config{
		App: AppConfig{
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		DB: DBConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnv("DB_PORT", "3306"),
			Name:            getEnv("DB_NAME", ""),
			User:            getEnv("DB_USER", ""),
			Password:        getEnv("DB_PASSWORD", ""),
			MaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 10),
			ConnMaxLifetime: getEnvDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
		},
		JWT: JWTConfig{
			Secret:            mustGetEnv("JWT_SECRET"),
			AccessExpiration:  mustGetDuration("JWT_ACCESS_EXPIRES_IN"),
			RefreshExpiration: getEnvDuration("JWT_REFRESH_EXPIRES_IN", 7*24*time.Hour), // Default 7 days
		},
	}
}

func Get() *Config {
	if cfg == nil {
		log.Fatal("config not loaded")
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Printf("invalid int for %s, using fallback: %v", key, fallback)
		return fallback
	}
	return i
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		log.Printf("invalid duration for %s, using fallback: %v", key, fallback)
		return fallback
	}
	return d
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required env: %s", key)
	}
	return v
}

func mustGetDuration(key string) time.Duration {
	v := mustGetEnv(key)
	d, err := time.ParseDuration(v)
	if err != nil {
		log.Fatalf("invalid duration for %s", key)
	}
	return d
}

func Run(server *http.Server) {

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
