package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/registry")
		userRouterPublic.POST("/otp")
	}
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/get-information")
		userRouterPrivate.POST("/otp")
	}
}
