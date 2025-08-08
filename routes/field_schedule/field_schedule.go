package routes

import (
	"field-service/clients"
	"field-service/constants"
	"field-service/controllers"
	"field-service/middlewares"
	"github.com/gin-gonic/gin"
)

type FieldScheduleRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	client     clients.IClientRegistry
}

type IFieldScheduleRoute interface {
	Run()
}

func NewFieldScheduleRoute(router *gin.Engine, controller controllers.IControllerRegistry, client clients.IClientRegistry) *FieldScheduleRoute {
	return &FieldScheduleRoute{
		controller: controller,
		group:      router.Group("/field"),
		client:     client,
	}
}

func (f *FieldScheduleRoute) Run() {
	group := f.group.Group("/field_schedule")
	group.GET("", middlewares.AuthenticateWithoutToken(), f.controller.GetFieldSchedule().GetAllByFieldIDAndDate)
	group.PATCH("", middlewares.AuthenticateWithoutToken(), f.controller.GetFieldSchedule().UpdateStatus)
	group.Use(middlewares.Authenticate())
	group.GET("/pagination", middlewares.CheckRole(
		[]string{
			constants.Admin,
			constants.Customer,
		}, f.client),
		f.controller.GetFieldSchedule().GetAllWithPagination,
	)
	group.GET("/:uuid", middlewares.CheckRole(
		[]string{
			constants.Admin,
			constants.Customer,
		}, f.client),
		f.controller.GetFieldSchedule().GetByUUID,
	)
	group.POST("", middlewares.CheckRole(
		[]string{
			constants.Admin,
		}, f.client),
		f.controller.GetFieldSchedule().Create,
	)
	group.POST("/one-month", middlewares.CheckRole(
		[]string{
			constants.Admin,
		}, f.client),
		f.controller.GetFieldSchedule().GenerateScheduleForOneMonth,
	)
	group.PUT("/:uuid", middlewares.CheckRole(
		[]string{
			constants.Admin,
		}, f.client),
		f.controller.GetFieldSchedule().Update,
	)
	group.DELETE("/:uuid", middlewares.CheckRole(
		[]string{
			constants.Admin,
		}, f.client),
		f.controller.GetFieldSchedule().Delete,
	)
}
