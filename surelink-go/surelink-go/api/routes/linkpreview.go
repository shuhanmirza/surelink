package routes

import (
	"surelink-go/api/controller"
	"surelink-go/infrastructure"
)

type LinkPreviewRoute struct {
	Controller controller.LinkPreviewController
	Handler    infrastructure.GinRouter
}

func NewLinkPreviewRoute(linkPreviewController controller.LinkPreviewController, router infrastructure.GinRouter) LinkPreviewRoute {
	return LinkPreviewRoute{
		Controller: linkPreviewController,
		Handler:    router,
	}
}

func (r *LinkPreviewRoute) Setup() {
	configuration := r.Handler.Gin.Group("link-preview")
	{
		configuration.GET("/", r.Controller.GetLinkPreview)
	}
}
