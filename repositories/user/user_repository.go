package user

import (
	"todo-app/dto/request"
	"todo-app/models"
)

type UserRepository interface {
	Register(request request.RegisterRequest) (models.User, error)
	FindByUsername(username string) (models.User, error)
	FindByEmail(email string) (models.User, error)
}
