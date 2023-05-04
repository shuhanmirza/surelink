package routes

import (
	"surelink-go/api/controller"
	"surelink-go/infrastructure"
)

type StatRoute struct {
	Controller controller.StatController
	Handler    infrastructure.GinRouter
}

func NewStatRoute(statController controller.StatController, router infrastructure.GinRouter) StatRoute {
	return StatRoute{
		Controller: statController,
		Handler:    router,
	}
}

func (r *StatRoute) Setup() {
	configuration := r.Handler.Gin.Group("stat")
	{
		configuration.GET("home", r.Controller.GetHomeStat)
	}

}
