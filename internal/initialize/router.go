package initialize

import (
	"github.com/binnguci/todo-app/global"
	"github.com/binnguci/todo-app/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	r.Use() //logging middleware
	r.Use() //cors middleware
	r.Use() //limiter middleware
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/api")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitTaskRouter(MainGroup)
	}
	{
		managerRouter.InitUserRouter(MainGroup)
		managerRouter.InitAdminRouter(MainGroup)
	}
	return r
}
