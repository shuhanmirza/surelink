package routes

import (
	"surelink-go/api/controller"
	"surelink-go/infrastructure"
)

type CaptchaRoute struct {
	Controller controller.CaptchaController
	Handler    infrastructure.GinRouter
}

func NewCaptchaRoute(captchaController controller.CaptchaController, router infrastructure.GinRouter) CaptchaRoute {
	return CaptchaRoute{
		Controller: captchaController,
		Handler:    router,
	}
}

func (r *CaptchaRoute) Setup() {
	configuration := r.Handler.Gin.Group("captcha")
	{
		configuration.GET("/new", r.Controller.GetCaptcha)
	}
}
