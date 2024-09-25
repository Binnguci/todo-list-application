package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/domain/response"
	"todo-app/error"
	"todo-app/services/tag"
)

type TagController struct {
	tagService tag.TagService
}

func NewTagController(tagService tag.TagService) *TagController {
	return &TagController{tagService: tagService}
}

func (t TagController) FindAll(ctx *gin.Context) {
	tag, err := t.tagService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  error.INTERNAL_SERVER_ERROR.Code,
			Message: error.INTERNAL_SERVER_ERROR.Message,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.APIResponse{
		Status:  error.SUCCESS.Code,
		Message: error.SUCCESS.Message,
		Data:    tag,
	})
}
