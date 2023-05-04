package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"surelink-go/api/service"
	"surelink-go/util"
)

type StatController struct {
	statService service.StatService
}

func NewStatController(statService service.StatService) StatController {
	return StatController{
		statService: statService,
	}
}

func (c *StatController) GetHomeStat(ctx *gin.Context) {
	response, err := c.statService.GetHomePageStats(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
