package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	IsCompleted bool           `json:"is_completed" gorm:"default:false"`
	Deadline    time.Time      `json:"deadline"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	User        User           `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
