package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/lakhan-purohit/net-http/internal/pkg/response"
)

type client struct {
	requests int
	lastSeen time.Time
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
)

const (
	maxRequests = 100
	window      = time.Minute
)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		mu.Lock()
		c, exists := clients[ip]
		if !exists {
			c = &client{}
			clients[ip] = c
		}

		if time.Since(c.lastSeen) > window {
			c.requests = 0
		}

		c.requests++
		c.lastSeen = time.Now()

		if c.requests > maxRequests {
			mu.Unlock()
			response.TooManyRequests(response.SendParams{
				W: w,
			})
			return
		}
		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
