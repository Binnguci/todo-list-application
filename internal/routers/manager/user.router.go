package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ar *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterPrivate := Router.Group("/admin/account")
	{
		userRouterPrivate.POST("/block-user")
	}
}
