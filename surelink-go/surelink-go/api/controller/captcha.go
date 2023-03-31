package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"surelink-go/api/service"
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
	ctx.JSON(http.StatusOK, nil)
}
