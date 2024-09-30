package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIResponse struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(http.StatusOK, APIResponse{
		code,
		msg[code],
		data,
	})
}

func ErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, APIResponse{
		code,
		msg[code],
		nil,
	})
}
