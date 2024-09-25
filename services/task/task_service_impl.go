package task

import (
	"todo-app/domain/request"
	"todo-app/models"
	"todo-app/repositories/task"
)

type TaskServiceImpl struct {
	taskRepository task.TaskRepository
}

func NewTaskServiceImpl(taskRepository task.TaskRepository) TaskService {
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

func (t *TaskServiceImpl) Create(newTask request.TaskRequest) (models.Task, error) {
	task, err := t.taskRepository.Create(newTask)
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
