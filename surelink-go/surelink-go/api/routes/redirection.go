package routes

import (
	"surelink-go/api/controller"
	"surelink-go/infrastructure"
)

type RedirectionRoute struct {
	Controller controller.RedirectionController
	Handler    infrastructure.GinRouter
}

func NewRedirectionRoute(redirectionController controller.RedirectionController, router infrastructure.GinRouter) RedirectionRoute {
	return RedirectionRoute{
		Controller: redirectionController,
		Handler:    router,
	}
}

func (r *RedirectionRoute) Setup() {
	configuration := r.Handler.Gin.Group("redirection")
	{
		configuration.GET("/get-map", r.Controller.GetMap)
		configuration.POST("/set-map", r.Controller.SetMap)
	}
}
