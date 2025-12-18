package model

import "time"

// UserStats represents complex data related to a user
// @Description User statistics and metadata
type UserStats struct {
	UserID     int64     `json:"user_id" db:"user_id"`
	LastLogin  time.Time `json:"last_login" db:"last_login"`
	LoginCount int       `json:"login_count" db:"login_count"`
}

// UserWithStats is a complex responded model
type UserWithStats struct {
	User
	Stats *UserStats `json:"stats"`
}
