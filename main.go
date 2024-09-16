package main

import (
	"log"
	"net/http"
	"time"
	"todo-app/config"
	"todo-app/controllers"
	"todo-app/repositories"
	"todo-app/routes"
	"todo-app/services"
)

func main() {
	config.ConnectDatabase()

	userRepository := repositories.NewUserRepositoryImpl(config.DB)
	userService := services.NewUserServiceImpl(userRepository)
	userController := controllers.NewUserController(userService)

	router := routes.UserRoutes(userController)
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
