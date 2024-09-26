package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"unique;not null"`
	Tasks     []Task         `json:"tasks" gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
