package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             uint           `json:"id" gorm:"primaryKey,autoIncrement"`
	Username       string         `json:"username" gorm:"unique;not null"`
	Email          string         `json:"email" gorm:"unique;not null"`
	Password       string         `json:"password" gorm:"not null"`
	FullName       string         `json:"full_name"`
	DateOfBirth    time.Time      `json:"date_of_birth"`
	ProfilePicture string         `json:"profile_picture"`
	OAuth2Provider string         `json:"oauth2_provider"`
	OAuth2ID       string         `json:"oauth2_id"`
	OAuth2Token    string         `json:"oauth2_token"`
	Status         string         `json:"status" gorm:"default:'active'"`
	Roles          []Role         `json:"roles" gorm:"many2many:user_roles"`
	LastLogin      time.Time      `json:"last_login"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
