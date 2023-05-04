package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"surelink-go/api/service"
	"surelink-go/api/structs"
	"surelink-go/util"
)

type LinkPreviewController struct {
	linkPreviewService service.LinkPreviewService
}

func NewLinkPreviewController(linkPreviewService service.LinkPreviewService) LinkPreviewController {
	return LinkPreviewController{
		linkPreviewService: linkPreviewService,
	}
}

func (c *LinkPreviewController) GetLinkPreview(ctx *gin.Context) {
	var request structs.GetLinkPreviewRequest

	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		util.HandleCommonValidationError(ctx, err)
		return
	}

	response, err := c.linkPreviewService.GetLinkPreview(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
