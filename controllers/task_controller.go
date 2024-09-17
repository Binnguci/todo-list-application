package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-app/dto/response"
	"todo-app/error"
	"todo-app/models"
	"todo-app/services"
)

type TaskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func (t *TaskController) FindAll(ctx *gin.Context) {
	tasks, err := t.taskService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  error.INTERNAL_SERVER_ERROR.Code,
			Message: error.INTERNAL_SERVER_ERROR.Message,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.APIResponse{
		Status:  error.SUCCESS.Code,
		Message: error.SUCCESS.Message,
		Data:    tasks,
	})
}

func (t *TaskController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	task, err := t.taskService.FindByID(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  error.INTERNAL_SERVER_ERROR.Code,
			Message: error.INTERNAL_SERVER_ERROR.Message,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.APIResponse{
		Status:  error.SUCCESS.Code,
		Message: error.SUCCESS.Message,
		Data:    task,
	})
}

func (t *TaskController) Create(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  error.INVALID_REQUEST.Code,
			Message: error.INVALID_REQUEST.Message,
		})
		return
	}

	taskResponse, err := t.taskService.Create(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  error.INTERNAL_SERVER_ERROR.Code,
			Message: error.INTERNAL_SERVER_ERROR.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.APIResponse{
		Status:  error.CREATE_SUCCESS.Code,
		Message: error.CREATE_SUCCESS.Message,
		Data:    taskResponse,
	})
}

func (t *TaskController) Update(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  error.INVALID_REQUEST.Code,
			Message: error.INVALID_REQUEST.Message,
		})
		return
	}

	taskResponse, err := t.taskService.Update(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  error.INTERNAL_SERVER_ERROR.Code,
			Message: error.INTERNAL_SERVER_ERROR.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.APIResponse{
		Status:  error.UPDATE_SUCCESS.Code,
		Message: error.UPDATE_SUCCESS.Message,
		Data:    taskResponse,
	})
}

func (t *TaskController) Delete(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  error.INVALID_REQUEST.Code,
			Message: error.INVALID_REQUEST.Message,
		})
		return
	}

	err := t.taskService.Delete(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  error.INTERNAL_SERVER_ERROR.Code,
			Message: error.INTERNAL_SERVER_ERROR.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.APIResponse{
		Status:  error.DELETE_SUCCESS.Code,
		Message: error.DELETE_SUCCESS.Message,
	})
}
