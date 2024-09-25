package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"unique;not null"`
	UserId    uint           `json:"user_id" gorm:"not null"`
	Tasks     []Task         `json:"tasks" gorm:"many2many:task_tags"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
