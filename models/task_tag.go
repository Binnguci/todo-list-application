package models

type TaskTag struct {
	TaskID uint `json:"task_id" gorm:"primaryKey"`
	TagID  uint `json:"tag_id" gorm:"primaryKey"`
}
