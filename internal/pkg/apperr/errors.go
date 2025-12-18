package apperr

import (
	"fmt"
	"net/http"
)

// AppError represents a professional application error
type AppError struct {
	Status  int    `json:"-"`
	Message string `json:"m"`    // Mapped to m in response
	Code    string `json:"c"`    // Mapped to c in response
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e.Status, e.Code, e.Message)
}

// Common errors
func New(status int, message string, code string) *AppError {
	return &AppError{
		Status:  status,
		Message: message,
		Code:    code,
	}
}

var (
	ErrNotFound      = New(http.StatusNotFound, "Resource not found", "NOT_FOUND")
	ErrUnauthorized  = New(http.StatusUnauthorized, "Unauthorized access", "UNAUTHORIZED")
	ErrBadRequest    = New(http.StatusStatusBadRequest, "Invalid request", "BAD_REQUEST")
	ErrInternal      = New(http.StatusInternalServerError, "Internal server error", "INTERNAL_ERROR")
)
