package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID             uuid.UUID  `json:"uuid" gorm:"type:char(36);primaryKey"`
	Username       string     `json:"username" gorm:"unique;not null"`
	Email          string     `json:"email" gorm:"unique;not null"`
	Password       string     `json:"password" gorm:"not null"`
	FullName       string     `json:"full_name"`
	DateOfBirth    *time.Time `json:"date_of_birth"`
	ProfilePicture string     `json:"profile_picture"`
	OAuth2Provider string     `json:"oauth2_provider"`
	OAuth2ID       string     `json:"oauth2_id"`
	OAuth2Token    string     `json:"oauth2_token"`
	IsActive       bool       `json:"is_active" gorm:"default:false"`
	Otp            int        `gorm:"column:otp"`
	LastLogin      *time.Time `json:"last_login"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func (u *User) TableName() string {
	return "users"
}
