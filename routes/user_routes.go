package routes

import (
	"github.com/gin-gonic/gin"
	"todo-app/controllers"
	"todo-app/middleware"
)

func UserRoutes(userController *controllers.UserController) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandler)
	router := r.Group("/api/user/")
	{
		router.POST("/register", userController.RegisterUser)
		router.POST("/login", userController.Login)
	}
	return r
}
