package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/dto/request"
	"todo-app/dto/response"
	"todo-app/error"
	"todo-app/services"
	"todo-app/utils"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}
func (u *UserController) RegisterUser(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid input",
			Error:   err.Error(),
		})
		return
	}

	user, err := u.userService.Register(req)
	if err != nil {
		if err == error.USERNAME_ALREADY_EXISTS || err == error.EMAIL_ALREADY_EXISTS {
			c.JSON(http.StatusConflict, response.APIResponse{
				Status:  http.StatusConflict,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to register user",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.APIResponse{
		Status:  http.StatusCreated,
		Message: "User registered successfully",
		Data:    user,
	})
}
func (u *UserController) Login(ctx *gin.Context) {
	var req request.AuthenticationRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := u.userService.VerifyUser(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Username, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	ctx.JSON(http.StatusOK, response.APIResponse{
		Status:  http.StatusOK,
		Message: "Login successful",
		Data:    gin.H{"token": token},
	})

}
