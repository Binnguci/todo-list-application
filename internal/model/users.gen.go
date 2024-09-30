// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	Username       string         `gorm:"column:username;not null" json:"username"`
	Email          string         `gorm:"column:email;not null" json:"email"`
	Password       string         `gorm:"column:password" json:"password"`
	Phone          string         `gorm:"column:phone" json:"phone"`
	FullName       string         `gorm:"column:full_name" json:"full_name"`
	Oauth2ID       string         `gorm:"column:oauth2_id" json:"oauth2_id"`
	Oauth2Provider string         `gorm:"column:oauth2_provider" json:"oauth2_provider"`
	Otp            string         `gorm:"column:otp" json:"otp"`
	OtpExpired     time.Time      `gorm:"column:otp_expired" json:"otp_expired"`
	RoleID         int32          `gorm:"column:role_id;not null" json:"role_id"`
	IsActive       bool           `gorm:"column:is_active" json:"is_active"`
	CreatedAt      time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
