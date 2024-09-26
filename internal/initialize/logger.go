package initialize

import (
	"github.com/binnguci/todo-app/global"
	"github.com/binnguci/todo-app/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
