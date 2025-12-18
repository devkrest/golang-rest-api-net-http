package model

// User represents a user in the system
// @Description User account information
type User struct {
	UUID         string `json:"uuid" db:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	ID           int64  `json:"id" db:"id" example:"1"`
	Username     string `json:"username" db:"username" example:"johndoe"`
	Email        string `json:"email" db:"email" example:"john@example.com"`
	Status       int    `json:"status" db:"status" example:"1"`
	Avatar       string `json:"avatar" db:"avatar" example:"avatar.jpg"`
	Token        string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string `json:"refresh_token" example:"def456..."`
}
