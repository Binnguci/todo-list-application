package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (a *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
}
