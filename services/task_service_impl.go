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

func (t *TaskServiceImpl) Update(id int, task models.Task) (models.Task, error) {
	updatedTask, err := t.taskRepository.Update(id, task)
	if err != nil {
		return models.Task{}, err
	}
	return updatedTask, nil
}

func (t *TaskServiceImpl) Delete(id int) error {
	err := t.taskRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
