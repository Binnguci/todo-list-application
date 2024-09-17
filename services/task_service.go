package services

import "todo-app/models"

type TaskService interface {
	FindAll() ([]models.Task, error)
	FindByID(id int) (models.Task, error)
	Create(task models.Task) (models.Task, error)
	Update(task models.Task) (models.Task, error)
	Delete(task models.Task) error
}
