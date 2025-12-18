# Golang API - Authentication & User Management

A robust, scalable REST API built with Go's standard `net/http` package. This project features JWT authentication, file uploads, MySQL integration, and modern API documentation using Scalar and Swagger.

## 🚀 Getting Started

### 1. Prerequisites
- **Go**: Download and install from [go.dev](https://go.dev/doc/install).
- **MySQL**: Ensure you have a MySQL server running.

### 2. Installation
Clone the repository and install dependencies:
```bash
git clone <repository-url>
cd net-http
go mod tidy
```

### 3. Environment Setup
Create a `.env` file in the root directory with the following content:
```env
# Server
APP_ENV=development
APP_PORT=8001

# MySQL
DB_HOST=localhost
DB_PORT=3306
DB_NAME=test
DB_USER=root
DB_PASSWORD=your_password

# Security
JWT_SECRET=your-secure-secret-key
JWT_EXPIRES_IN=24h
```

---

## 🛠️ Development Tools

### 1. Air (Live Reload)
`air` is used for hot-reloading the server during development.
- **Install**: `go install github.com/air-verse/air@latest`
- **Run**: Simply type `air` in the root directory.

### 2. Swagger (Documentation Generator)
`swag` generates the OpenAPI specification from code annotations.
- **Install**: `go install github.com/swaggo/swag/cmd/swag@latest`
- **Generate**: `swag init -g cmd/main.go`
- **Access**: [http://localhost:8001/swagger/index.html](http://localhost:8001/swagger/index.html)

### 3. Scalar (Modern API UI)
A beautiful, high-performance alternative to Swagger UI.
- **Access**: [http://localhost:8001/scalar](http://localhost:8001/scalar)
- **Features**: Includes a built-in API client, models sidebar, and search.

---

## 📂 Project Structure
- `cmd/`: Entry point of the application.
- `internal/`: Core logic (handlers, services, repositories, middleware).
- `docs/`: Generated Swagger and Scalar documentation files.
- `internal/pkg/`: Shared utilities and packages.

---

## 📜 Master Makefile (Elite Automation)
The project includes a `Makefile` for one-command developer automation:
- **`make run`**: Fire up the server with Air (Live Reload).
- **`make docs`**: Regenerate both Swagger and Scalar documentation.
- **`make tidy`**: Clean up and update Go modules.
- **`make build`**: Compile the project into a professional production binary.

---

## 🏎️ Performance & Security (Built-in)
- **Gzip Compression**: Automatically shrinks JSON responses for faster delivery.
- **Request Tracing**: Uses `X-Request-ID` to track logs across services.
- **Monitoring**: Every response includes `X-Response-Time` in headers.
- **Security**: Professional headers (`HSTS`, `CSP`, `X-Frame`) configured by default.
- **Optimized SQL**: Reflective database scanner with plan caching for near-native speed.

## 🛡️ License
Distributed under the Apache 2.0 License. See `LICENSE` for more information.
