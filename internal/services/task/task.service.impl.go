package task

import (
	"todo-app/domain/request"
	"todo-app/internal/models"
	"todo-app/internal/repositories/task"
)

type TaskServiceImpl struct {
	taskRepository task.TaskRepository
}

func NewTaskServiceImpl(taskRepository task.TaskRepository) TaskService {
	return &TaskServiceImpl{taskRepository: taskRepository}
}

func (tsi *TaskServiceImpl) FindAll() ([]models.Task, error) {
	task, err := tsi.taskRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (tsi *TaskServiceImpl) FindByID(id int) (models.Task, error) {
	task, err := tsi.taskRepository.FindByID(id)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (tsi *TaskServiceImpl) Create(newTask request.TaskRequest) (models.Task, error) {
	task, err := tsi.taskRepository.Create(newTask)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (tsi *TaskServiceImpl) Update(id int, task models.Task) (models.Task, error) {
	updatedTask, err := tsi.taskRepository.Update(id, task)
	if err != nil {
		return models.Task{}, err
	}
	return updatedTask, nil
}

func (tsi *TaskServiceImpl) Delete(id int) error {
	err := tsi.taskRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
