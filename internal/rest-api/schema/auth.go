package schema

import "mime/multipart"

type LoginRequest struct {
	Email    string `form:"email" json:"email" validate:"required,email" example:"john@example.com"`
	Password string `form:"password" json:"password" validate:"required,min=6" example:"password123"`
}

type SignUpRequest struct {
	Username string `json:"username" form:"username" validate:"required,min=3,max=30" example:"johndoe"`
	Email    string `json:"email" form:"email" validate:"required,email" example:"john@example.com"`
	Password string `json:"password" form:"password" validate:"required,min=6" example:"password123"`

	Avatar       *multipart.FileHeader `file:"avatar"`
	LicenseFront *multipart.FileHeader `file:"license_front"`
	LicenseBack  *multipart.FileHeader `file:"license_back"`
}
