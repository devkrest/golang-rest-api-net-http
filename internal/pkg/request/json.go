package request

import (
	"encoding/json"
	"net/http"

	"github.com/lakhan-purohit/net-http/internal/pkg/response"
)

func JSON(w http.ResponseWriter, r *http.Request, dst any) bool {
	if r.Body == nil {
		response.BadRequest(response.SendParams{
			W:       w,
			Message: "request body required",
		})
		return false
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // ✅ strict JSON

	if err := decoder.Decode(dst); err != nil {
		response.BadRequest(response.SendParams{
			W:       w,
			Message: "invalid json body",
		})
		return false
	}

	return true
}
