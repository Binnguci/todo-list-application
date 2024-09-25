package task

import (
	"todo-app/domain/request"
	"todo-app/models"
)

type TaskRepository interface {
	FindAll() ([]models.Task, error)
	FindByID(id int) (models.Task, error)
	Create(task request.TaskRequest) (models.Task, error)
	Update(id int, task models.Task) (models.Task, error)
	Delete(id int) error
}
