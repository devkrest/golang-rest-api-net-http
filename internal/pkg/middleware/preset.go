package middleware

import "net/http"

func Public(h http.Handler) http.Handler {
	return Chain(
		h,
		Recover,
		Gzip,
		SecurityHeaders,
		RequestID,
		CORS,
		RateLimit,
		Timer,
		Logger,
	)
}

func Private(h http.Handler) http.Handler {
	return Chain(
		h,
		Recover,
		Gzip,
		SecurityHeaders,
		RequestID,
		CORS,
		JWT,
		RateLimit,
		Timer,
		Logger,
	)
}

func Protected(h http.Handler) http.Handler {
	return Chain(
		h,
		APIKey,
	)
}
