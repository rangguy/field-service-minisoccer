package routes

import (
	"field-service/controllers"
	routes "field-service/routes/user"
	"github.com/gin-gonic/gin"
)

type Registry struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type IRouterRegistry interface {
	Serve()
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup) *Registry {
	return &Registry{
		controller: controller,
		group:      group,
	}
}

func (r *Registry) Serve() {
	r.userRoute().Run()
}

func (r *Registry) userRoute() routes.IUserRoute {
	return routes.NewUserRoute(r.controller, r.group)
}
