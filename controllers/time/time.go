package controllers

import (
	errValidation "field-service/common/error"
	"field-service/common/response"
	"field-service/domain/dto"
	"field-service/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func NewTimeController(service services.IServiceRegistry) ITimeController {
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
	err := context.ShouldBind(&request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusBadRequest,
			Err:     err,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     context,
		})
		return
	}

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
