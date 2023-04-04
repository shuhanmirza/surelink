package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"surelink-go/api/service"
	"surelink-go/util"
)

type CaptchaController struct {
	captchaService service.CaptchaService
}

func NewCaptchaController(captchaService service.CaptchaService) CaptchaController {
	return CaptchaController{
		captchaService: captchaService,
	}
}

func (c CaptchaController) GetCaptcha(ctx *gin.Context) {
	getCaptchaResponse, err := c.captchaService.GetNewCaptcha(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, getCaptchaResponse)
}
