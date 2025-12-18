package middleware

	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/lakhan-purohit/net-http/internal/pkg/constants"
	"github.com/lakhan-purohit/net-http/internal/pkg/response"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				requestID, _ := r.Context().Value(constants.RequestIDContextKey).(string)

				slog.Error("panic_recovered",
					"error", err,
					"request_id", requestID,
					"stack", string(debug.Stack()),
				)

				response.InternalError(response.SendParams{
					W:       w,
					Message: "A serious error occurred. Please contact support.",
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
