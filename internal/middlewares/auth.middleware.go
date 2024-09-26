package middlewares

import (
	"github.com/binnguci/todo-app/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			response.ErrorResponse(context, response.ErrInvalidToken, "")
			context.Abort()
			return
		}
		context.Next()
	}
}
