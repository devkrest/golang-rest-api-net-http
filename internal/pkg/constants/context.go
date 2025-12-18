package constants

type CtxKey string

const (
	UserContextKey      CtxKey = "user_claims"
	RequestIDContextKey CtxKey = "request_id"
)
