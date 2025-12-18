# 🚀 Golang API - Elite Go REST Standard

[![Go Version](https://img.shields.io/github/go-mod/go-version/lekhnath/net-http?color=00ADD8&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Standard](https://img.shields.io/badge/standard-net%2Fhttp-gold)](https://pkg.go.dev/net/http)

A high-performance, production-ready REST API built with Go's standard `net/http` package. This project serves as a "Golden Standard" for Go API development, featuring a robust middleware stack, type-safe utilities, and professional automation.

---

## 🏎️ Elite Features & Performance
- **Zero-Dependency Core**: Built almost entirely on `net/http` for maximum longevity.
- **High-Speed Database**: Reflective scanner with **execution plan caching** for near-native performance.
- **Resilient Middleware**: A comprehensive stack for security, observability, and performance.
- **Modern Docs**: Dual support for **Swagger** and **Scalar** (beautiful, modern UI).
- **Type-Safe JWT**: Advanced JWT handling with separate access and refresh token life-cycles.

---

## 🛠️ Getting Started

### 1. Prerequisites
- **Go**: 1.23+ recommended.
- **MySQL**: 8.0+ for optimal compatibility.

### 2. Installation
```bash
git clone <repository-url>
cd net-http
go mod tidy
```

### 3. Database Setup
1. Create a MySQL database (e.g., `golang_api`).
2. Run the initial schema:
```bash
mysql -u root -p golang_api < schema.sql
```

### 4. Environment Setup
Update your `.env` file with these professional-grade settings (see `.env.example` for reference):
```env
# Server Configuration
APP_ENV=development
APP_PORT=8001

# MySQL Database (Connection Pooling Enabled)
DB_HOST=localhost
DB_PORT=3306
DB_NAME=golang_api
DB_USER=root
DB_PASSWORD=
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=10
DB_CONN_MAX_LIFETIME=5m

# Security & JWT
JWT_SECRET=your-secure-secret-key
JWT_ACCESS_EXPIRES_IN=1h
JWT_REFRESH_EXPIRES_IN=168h
```

---

## 📜 Master Makefile (Elite Automation)
One command to rule them all. High-efficiency developer workflows:
- `make run`: Starts development with **Air (Live Reload)**.
- `make docs`: Regenerates both **Swagger** and **Scalar** documentation.
- `make tidy`: Automates module cleanup and security updates.
- `make build`: Produces an optimized, production-ready binary.

---

## 🛡️ Elite Middleware Stack
Every request passes through a meticulously crafted pipeline:
| Middleware | Function |
| :--- | :--- |
| **Request ID** | Injects `X-Request-ID` for end-to-end tracing. |
| **Panic Recovery** | Gracefully handles crashes with full stack-trace logging via `slog`. |
| **Security Headers** | Enforces `HSTS`, `CSP`, `XSS-Protection`, and `Frame-Options`. |
| **Timer** | Injects `X-Response-Time` to monitor API latency. |
| **Gzip** | Transparent JSON compression for bandwidth optimization. |
| **Rate Limiter** | Prevents abuse through sophisticated request throttling. |
| **CORS** | Securely handles cross-origin requests for frontend integration. |

---

## 📂 Project Architecture
```text
├── cmd/               # Application entry point
├── docs/              # Swagger & Scalar documentation specs
├── internal/
│   ├── rest-api/      # domain logic (handlers, repositories, services)
│   └── pkg/           # High-performance internal packages
│       ├── middleware/ # Elite middleware stack
│       ├── db/         # Database engine & scanner
│       └── utils/      # Type-safe crypto, JWT, and file utils
└── .air.toml          # Hot-reload configuration
```

---

## 📖 API Documentation
- **Scalar UI (Recommended)**: [http://localhost:8001/scalar](http://localhost:8001/scalar)
- **Swagger UI**: [http://localhost:8001/swagger/index.html](http://localhost:8001/swagger/index.html)

---

## 🛡️ License
Distributed under the Apache 2.0 License.
