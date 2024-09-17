package repositories

import (
	"gorm.io/gorm"
	"todo-app/models"
)

type TaskRepositoryImpl struct {
	Db *gorm.DB
}

func NewTaskRepositoryImpl(db *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{Db: db}
}

func (t *TaskRepositoryImpl) FindAll() ([]models.Task, error) {
	err := t.Db.Find(&[]models.Task{}).Error
	if err != nil {
		return nil, err
	}
	return []models.Task{}, nil
}

func (t *TaskRepositoryImpl) FindByID(id int) (models.Task, error) {
	err := t.Db.First(&models.Task{}, id).Error
	if err != nil {
		return models.Task{}, err
	}
	return models.Task{}, nil
}

func (t *TaskRepositoryImpl) Create(task models.Task) (models.Task, error) {
	err := t.Db.Create(&task).Error
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (t *TaskRepositoryImpl) Update(task models.Task) (models.Task, error) {
	err := t.Db.Save(&task).Error
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (t *TaskRepositoryImpl) Delete(task models.Task) error {
	err := t.Db.Delete(&task).Error
	if err != nil {
		return err
	}
	return nil
}
