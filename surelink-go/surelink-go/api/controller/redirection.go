package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"surelink-go/api/service"
	"surelink-go/api/structs"
	"surelink-go/util"
)

type RedirectionController struct {
	redirectionService service.RedirectionService
}

func NewRedirectionController(redirectionService service.RedirectionService) RedirectionController {
	return RedirectionController{
		redirectionService: redirectionService,
	}
}

func (c *RedirectionController) GetMap(ctx *gin.Context) {
	var request structs.GetMapRequest

	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		util.HandleCommonValidationError(ctx, err)
		return
	}

	response, err := c.redirectionService.GetMap(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *RedirectionController) SetMap(ctx *gin.Context) {
	var request structs.SetMapRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.HandleCommonValidationError(ctx, err)
		return
	}

	response, err := c.redirectionService.SetMap(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
