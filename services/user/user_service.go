package user

import (
	"todo-app/domain/request"
	"todo-app/domain/response"
)

type UserService interface {
	Register(request request.RegisterRequest) (response.UserResponse, error)
	VerifyUser(username string, password string) (response.UserResponse, error)
}
