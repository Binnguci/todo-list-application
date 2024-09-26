package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             uuid.UUID      `json:"id" gorm:"type:char(36);default:uuid()"`
	Username       string         `json:"username" gorm:"unique;not null"`
	Email          string         `json:"email" gorm:"unique;not null"`
	Password       string         `json:"password" gorm:"not null"`
	FullName       string         `json:"full_name"`
	DateOfBirth    *time.Time     `json:"date_of_birth"`
	ProfilePicture string         `json:"profile_picture"`
	OAuth2Provider string         `json:"oauth2_provider" gorm:"column:oauth2_provider"`
	OAuth2ID       string         `json:"oauth2_id" gorm:"column:oauth2_id"`
	OAuth2Token    string         `json:"oauth2_token" gorm:"column:oauth2_token"`
	Status         string         `json:"status" gorm:"default:'off'"`
	Otp            int            `gorm:"otp"`
	LastLogin      *time.Time     `json:"last_login"`
	Tasks          []Task         `json:"tasks" gorm:"foreignKey:UserID"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
