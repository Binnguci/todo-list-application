package services

import (
	"todo-app/models"
	"todo-app/repositories"
)

type TaskServiceImpl struct {
	taskRepository repositories.TaskRepository
}

func NewTaskServiceImpl(taskRepository repositories.TaskRepository) TaskService {
	return &TaskServiceImpl{taskRepository: taskRepository}
}

func (t *TaskServiceImpl) FindAll() ([]models.Task, error) {
	task, err := t.taskRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskServiceImpl) FindByID(id int) (models.Task, error) {
	task, err := t.taskRepository.FindByID(id)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (t *TaskServiceImpl) Create(task models.Task) (models.Task, error) {
	task, err := t.taskRepository.Create(task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (t *TaskServiceImpl) Update(task models.Task) (models.Task, error) {
	task, err := t.taskRepository.Update(task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (t *TaskServiceImpl) Delete(task models.Task) error {
	err := t.taskRepository.Delete(task)
	if err != nil {
		return err
	}
	return nil
}
