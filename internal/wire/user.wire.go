//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/binnguci/todo-app/internal/controllers"
	uri "github.com/binnguci/todo-app/internal/repositories/user"
	usi "github.com/binnguci/todo-app/internal/services/user"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUserRouterHandle(DB *gorm.DB) (*controllers.UserController, error) {
	wire.Build(
		uri.NewUserRepositoryImpl,
		usi.NewUserServiceImpl,
		controllers.NewUserController,
	)
	return new(controllers.UserController), nil
}
