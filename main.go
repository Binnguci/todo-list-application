package main

import (
	"log"
	"net/http"
	"time"
	"todo-app/config"
	"todo-app/controllers"
	"todo-app/repositories/task"
	"todo-app/repositories/user"
	"todo-app/routes"
	task2 "todo-app/services/task"
	user2 "todo-app/services/user"
)

func main() {
	config.ConnectDatabase()

	userRepository := user.NewUserRepositoryImpl(config.DB)
	userService := user2.NewUserServiceImpl(userRepository)
	userController := controllers.NewUserController(userService)

	taskRepository := task.NewTaskRepositoryImpl(config.DB)
	taskService := task2.NewTaskServiceImpl(taskRepository)
	taskController := controllers.NewTaskController(taskService)

	router := routes.UserRoutes(userController, taskController)
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
