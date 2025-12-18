package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/lakhan-purohit/net-http/internal/pkg/response"
	"github.com/lakhan-purohit/net-http/internal/pkg/utils"
)

type TokenPayload struct {
	DateTime string `json:"date_time"`
}

func APIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1️⃣ Bypass routes
		if strings.Contains(r.URL.Path, "webhook") ||
			strings.Contains(r.URL.Path, "swagger") {
			next.ServeHTTP(w, r)
			return
		}

		// 2️⃣ Read header
		key := r.Header.Get("x-api-key")
		if key == "" {
			response.TokenMissing(response.SendParams{
				W: w,
			})
			return
		}

		// 3️⃣ Dev bypass
		if key == "!!Devkrest!!" {
			next.ServeHTTP(w, r)
			return
		}

		// 4️⃣ Decrypt key
		decrypted, err := utils.DecryptKey(key)
		if err != nil {
			response.UnauthorizedAccess(response.SendParams{
				W: w,
			})
			return
		}

		// 5️⃣ Parse JSON
		var payload TokenPayload
		if err := json.Unmarshal([]byte(decrypted), &payload); err != nil {
			response.UnauthorizedAccess(response.SendParams{
				W: w,
			})
			return
		}

		// 6️⃣ Validate time
		tokenTime, err := time.Parse(time.RFC3339, payload.DateTime)
		if err != nil {
			response.UnauthorizedAccess(response.SendParams{
				W: w,
			})
			return
		}

		if time.Since(tokenTime) > time.Minute {
			response.UnauthorizedAccess(response.SendParams{
				W: w,
			})
			return
		}

		// ✅ Success
		next.ServeHTTP(w, r)
	})
}
