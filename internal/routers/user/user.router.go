package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userRouterPublic := Router.Group("/account")
	{
		userRouterPublic.POST("/registry")
		userRouterPublic.POST("/otp")
	}
	userRouterPrivate := Router.Group("/account")
	{
		userRouterPrivate.GET("/get-information")
	}
}
