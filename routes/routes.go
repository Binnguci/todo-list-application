package routes

import (
	"github.com/gin-gonic/gin"
	"todo-app/controllers"
	"todo-app/middleware"
)

func UserRoutes(userController *controllers.UserController, taskController *controllers.TaskController) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandler)
	userRoutes := r.Group("/api/user/")
	{
		userRoutes.POST("/register", userController.RegisterUser)
		userRoutes.POST("/login", userController.Login)
	}
	taskRoutes := r.Group("/api/task")
	{
		taskRoutes.GET("/", taskController.FindAll)
		taskRoutes.GET("/:id", taskController.FindById)
		taskRoutes.POST("/", taskController.Create)
		taskRoutes.PUT("/:id", taskController.Update)
		taskRoutes.DELETE("/:id", taskController.Delete)
	}
	r.Use(middleware.ErrorHandlingMiddleware)
	return r
}
