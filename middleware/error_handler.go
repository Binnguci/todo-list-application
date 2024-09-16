package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"todo-app/dto/response"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors[0].Err

		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})

		fmt.Printf("Error: %v\n", err)
	}
}
