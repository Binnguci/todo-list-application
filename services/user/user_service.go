package user

import (
	"todo-app/dto/request"
	"todo-app/dto/response"
)

type UserService interface {
	Register(request request.RegisterRequest) (response.UserResponse, error)
	VerifyUser(username string, password string) (response.UserResponse, error)
}
