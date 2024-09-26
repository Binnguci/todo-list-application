package routers

import (
	"github.com/binnguci/todo-app/internal/routers/manager"
	"github.com/binnguci/todo-app/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
