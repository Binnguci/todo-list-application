package services

import "todo-app/models"

type TaskService interface {
	FindAll() ([]models.Task, error)
	FindByID(id int) (models.Task, error)
	Create(task models.Task) (models.Task, error)
	Update(id int, task models.Task) (models.Task, error)
	Delete(id int) error
}
