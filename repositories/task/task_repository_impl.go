package task

import (
	"gorm.io/gorm"
	"todo-app/dto/request"
	"todo-app/models"
)

type TaskRepositoryImpl struct {
	Db *gorm.DB
}

func NewTaskRepositoryImpl(db *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{Db: db}
}

func (t *TaskRepositoryImpl) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	err := t.Db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskRepositoryImpl) FindByID(id int) (models.Task, error) {
	err := t.Db.First(&models.Task{}, id).Error
	if err != nil {
		return models.Task{}, err
	}
	return models.Task{}, nil
}

func (t *TaskRepositoryImpl) Create(newTask request.TaskRequest) (models.Task, error) {
	task := models.Task{
		Title:       newTask.Title,
		Description: newTask.Description,
		UserID:      newTask.UserID,
	}

	err := t.Db.Create(&task).Error
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (t *TaskRepositoryImpl) Update(id int, task models.Task) (models.Task, error) {
	var existingTask models.Task
	if err := t.Db.First(&existingTask, id).Error; err != nil {
		return models.Task{}, err
	}

	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.IsCompleted = task.IsCompleted
	existingTask.Deadline = task.Deadline

	if err := t.Db.Save(&existingTask).Error; err != nil {
		return models.Task{}, err
	}

	return existingTask, nil
}

func (t *TaskRepositoryImpl) Delete(id int) error {
	var task models.Task
	if err := t.Db.First(&task, id).Error; err != nil {
		return err
	}
	return t.Db.Delete(&task).Error
}
