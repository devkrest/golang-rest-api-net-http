package response

import (
	"encoding/json"
	"net/http"

	"github.com/lakhan-purohit/net-http/internal/pkg/apperr"
	"github.com/lakhan-purohit/net-http/internal/rest-api/model"
)

// SendParams defines input for sending a response
type SendParams struct {
	W       http.ResponseWriter
	Status  int
	Data    any
	Message string
}

// Response is the standard API response structure
type Response struct {
	Status  int    `json:"s" example:"1"`          // 1 for success, 0 for error
	Message string `json:"m" example:"Success"`    // Descriptive message
	Result  any    `json:"r,omitempty"`            // Response data
	Code    string `json:"c,omitempty" example:""` // Optional error code
}

// SuccessResponse is for Swagger documentation
// @Description Successful response structure
type SuccessResponse struct {
	Status  int    `json:"s" example:"1"`
	Message string `json:"m" example:"Success"`
	Result  any    `json:"r,omitempty"`
}

// LoginResponse is for Swagger documentation
// @Description Successful login response
type LoginResponse struct {
	Status  int        `json:"s" example:"1"`
	Message string     `json:"m" example:"Success"`
	Result  model.User `json:"r"`
}

// UserListResponse is for Swagger documentation
// @Description Successful user list response
type UserListResponse struct {
	Status  int            `json:"s" example:"1"`
	Message string         `json:"m" example:"Success"`
	Result  []model.User `json:"r"`
}

// UserFullListResponse is for complex data examples
// @Description Successful user list with stats response
type UserFullListResponse struct {
	Status  int                   `json:"s" example:"1"`
	Message string                `json:"m" example:"Success"`
	Result  []model.UserWithStats `json:"r"`
}

// ErrorResponse is for Swagger documentation
// @Description Error response structure
type ErrorResponse struct {
	Status  int    `json:"s" example:"0"`
	Message string `json:"m" example:"Error message"`
}

// send is the core response writer (DRY, internal use only)
func send(res SendParams, success bool, defaultStatus int, defaultMessage string) {
	w := res.W

	// Resolve HTTP status
	status := res.Status
	if status == 0 {
		status = defaultStatus
	}

	// Resolve message
	message := res.Message
	if message == "" {
		message = defaultMessage
	}

	payload := Response{
		Message: message,
		Result:  res.Data,
	}

	if success {
		payload.Status = 1
	} else {
		payload.Status = 0
		// Check if data is an AppError to extract the code
		if ae, ok := res.Data.(*apperr.AppError); ok {
			payload.Code = ae.Code
			payload.Result = nil // Don't send error object in Result field
		}
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// ============================
// Public helpers
// ============================

// Success sends a success response (200 OK)
func Success(res SendParams) {
	send(res, true, http.StatusOK, "Success")
}

// Created sends a success response (201 Created)
func Created(res SendParams) {
	send(res, true, http.StatusCreated, "Created")
}

// BadRequest sends a 400 response
func BadRequest(res SendParams) {
	send(res, false, http.StatusBadRequest, "Bad Request")
}

// Unauthorized sends a 401 response
func Unauthorized(res SendParams) {
	send(res, false, http.StatusUnauthorized, "Unauthenticated")
}

// UnauthorizedAccess sends a 401 response (Alias for Unauthorized)
func UnauthorizedAccess(res SendParams) {
	send(res, false, http.StatusUnauthorized, "Unauthorized Access")
}

// Forbidden sends a 403 response
func Forbidden(res SendParams) {
	send(res, false, http.StatusForbidden, "Forbidden")
}

// NotFound sends a 404 response
func NotFound(res SendParams) {
	send(res, false, http.StatusNotFound, "Not Found")
}

// InternalError sends a 500 response
func InternalError(res SendParams) {
	send(res, false, http.StatusInternalServerError, "Internal Server Error")
}

// TooManyRequests sends a 429 response
func TooManyRequests(res SendParams) {
	send(res, false, http.StatusTooManyRequests, "Too Many Requests")
}

// MethodNotAllowed sends a 405 response
func MethodNotAllowed(res SendParams) {
	send(res, false, http.StatusMethodNotAllowed, "Method Not Allowed")
}

// TOKEN_MISSING
func TokenMissing(res SendParams) {
	send(res, false, http.StatusUnauthorized, "Token Missing")
}

// Error handles a professional application error
func Error(w http.ResponseWriter, err error) {
	if ae, ok := err.(*apperr.AppError); ok {
		send(SendParams{
			W:       w,
			Status:  ae.Status,
			Message: ae.Message,
			Data:    ae,
		}, false, ae.Status, ae.Message)
		return
	}

	// Default to internal error if not an AppError
	InternalError(SendParams{W: w, Message: err.Error()})
}
