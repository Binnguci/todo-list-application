package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ar *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterPrivate := Router.Group("/admin/user")
	{
		userRouterPrivate.POST("/block-user")
	}
}
