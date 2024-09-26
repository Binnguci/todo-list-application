package user

import "github.com/gin-gonic/gin"

type TaskRouter struct{}

func (tr *TaskRouter) InitTaskRouter(Router *gin.RouterGroup) {
	taskRouterPublic := Router.Group("/task")
	{
		taskRouterPublic.GET("/search")
		taskRouterPublic.GET("/list")
		taskRouterPublic.GET(":id")
	}
}
