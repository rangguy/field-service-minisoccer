package controllers

import (
	"field-service/common/response"
	"field-service/domain/dto"
	"field-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TimeController struct {
	service services.IServiceRegistry
}

type ITimeController interface {
	GetAll(*gin.Context)
	GetByUUID(*gin.Context)
	Create(*gin.Context)
}

func NewTimeController(service services.IServiceRegistry) *TimeController {
	return &TimeController{service: service}
}

func (t *TimeController) GetAll(context *gin.Context) {
	result, err := t.service.GetTime().GetAll(context)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  context,
	})
}

func (t *TimeController) GetByUUID(context *gin.Context) {
	uuid := context.Param("uuid")
	result, err := t.service.GetTime().GetByUUID(context, uuid)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  context,
	})
}

func (t *TimeController) Create(context *gin.Context) {
	var request dto.TimeRequest
	result, err := t.service.GetTime().Create(context, &request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusCreated,
		Data: result,
		Gin:  context,
	})
}
