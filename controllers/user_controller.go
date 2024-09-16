package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/dto/request"
	"todo-app/dto/response"
	"todo-app/error"
	"todo-app/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var req request.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  error.INVALID_REQUEST.Code,
			Message: error.INVALID_REQUEST.Message,
		})
		return
	}

	userResponse, err := c.userService.Register(req)
	if err != nil {
		var appErr error.AppError
		if errors.As(err, &appErr) {
			ctx.JSON(appErr.Code, response.APIResponse{
				Status:  appErr.Code,
				Message: appErr.Message,
			})
		} else {
			ctx.JSON(error.INTERNAL_SERVER_ERROR.Code, response.APIResponse{
				Status:  error.INTERNAL_SERVER_ERROR.Code,
				Message: error.INTERNAL_SERVER_ERROR.Message,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, response.APIResponse{
		Status:  error.CREATE_SUCCESS.Code,
		Message: error.CREATE_SUCCESS.Message,
		Data:    userResponse,
	})
}
