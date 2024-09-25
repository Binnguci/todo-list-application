package routes

import (
	"github.com/gin-gonic/gin"
	"todo-app/controllers"
	"todo-app/middleware"
)

type RouteControllers struct {
	UserController *controllers.UserController
	TaskController *controllers.TaskController
	TagController  *controllers.TagController
}

func InitRoutes(controllers RouteControllers) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.ErrorHandler)

	initUserRoutes(r, controllers.UserController)
	initTaskRoutes(r, controllers.TaskController)
	initTagRoutes(r, controllers.TagController)

	r.Use(middleware.ErrorHandlingMiddleware)
	return r
}

func initUserRoutes(r *gin.Engine, userController *controllers.UserController) {
	userRoutes := r.Group("/api/user")
	{
		userRoutes.POST("/register", userController.RegisterUser)
		userRoutes.POST("/login", userController.Login)
	}
}

func initTaskRoutes(r *gin.Engine, taskController *controllers.TaskController) {
	taskRoutes := r.Group("/api/task")
	{
		taskRoutes.GET("/", taskController.FindAll)
		taskRoutes.GET("/:id", taskController.FindById)
		taskRoutes.POST("/", taskController.Create)
		taskRoutes.PUT("/:id", taskController.Update)
		taskRoutes.DELETE("/:id", taskController.Delete)
	}
}

func initTagRoutes(r *gin.Engine, tagController *controllers.TagController) {
	tagRoutes := r.Group("/api/tag")
	{
		tagRoutes.GET("/", tagController.FindAll)
	}
}
