package main

import (
	"log"
	"net/http"
	"time"
	"todo-app/config"
	"todo-app/controllers"
	"todo-app/repositories/tag"
	"todo-app/repositories/task"
	"todo-app/repositories/user"
	"todo-app/routes"
	tag2 "todo-app/services/tag"
	task2 "todo-app/services/task"
	user2 "todo-app/services/user"
)

type Controllers struct {
	UserController *controllers.UserController
	TaskController *controllers.TaskController
	TagController  *controllers.TagController
}

func main() {
	config.ConnectDatabase()

	userRepository := user.NewUserRepositoryImpl(config.DB)
	userService := user2.NewUserServiceImpl(userRepository)
	userController := controllers.NewUserController(userService)

	taskRepository := task.NewTaskRepositoryImpl(config.DB)
	taskService := task2.NewTaskServiceImpl(taskRepository)
	taskController := controllers.NewTaskController(taskService)

	tagRepository := tag.NewTagRepositoryImpl(config.DB)
	tagService := tag2.NewTagServiceImpl(tagRepository)
	tagController := controllers.NewTagController(tagService)

	ctrls := Controllers{
		UserController: userController,
		TaskController: taskController,
		TagController:  tagController,
	}
	routeControllers := routes.RouteControllers{
		UserController: ctrls.UserController,
		TaskController: ctrls.TaskController,
		TagController:  ctrls.TagController,
	}

	router := routes.InitRoutes(routeControllers)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
