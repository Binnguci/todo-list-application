package response

import (
	"gorm.io/gorm"
	"time"
)

type UserResponse struct {
	ID             string         `json:"id"`
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	Phone          string         `json:"phone"`
	FullName       string         `json:"full_name"`
	Oauth2Provider string         `json:"oauth2_provider"`
	OtpExpired     time.Time      `json:"otp_expired"`
	RoleID         int32          `json:"role_id"`
	IsActive       bool           `json:"is_active"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}
