package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"todo-app/domain/response"
)

func ErrorHandler(c *gin.Context) {
	if len(c.Errors) > 0 {
		for _, e := range c.Errors {
			err := e.Err

			if c.Writer.Status() == http.StatusOK {
				c.JSON(http.StatusInternalServerError, response.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "Internal Server Error",
					Error:   err.Error(),
				})
			} else {
				c.JSON(c.Writer.Status(), response.APIResponse{
					Status:  c.Writer.Status(),
					Message: http.StatusText(c.Writer.Status()),
					Error:   err.Error(),
				})
			}
			fmt.Printf("Error: %v\n", err)
		}
	}
	c.Next()
}

func ErrorHandlingMiddleware(ctx *gin.Context) {
	fmt.Print("ErrorHandlingMiddleware\n ", ctx.Errors)
	ctx.Next()
}
