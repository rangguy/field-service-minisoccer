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

type FieldScheduleController struct {
	service services.IServiceRegistry
}

type IFieldScheduleController interface {
	GetAllWithPagination(*gin.Context)
	GetAllByFieldIDAndDate(*gin.Context)
	GetByUUID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	UpdateStatus(*gin.Context)
	Delete(*gin.Context)
	GenerateScheduleForOneMonth(*gin.Context)
}

func NewFieldScheduleController(service services.IServiceRegistry) IFieldScheduleController {
	return &FieldScheduleController{
		service: service,
	}
}

func (f *FieldScheduleController) GetAllWithPagination(context *gin.Context) {
	var params dto.FieldScheduleRequestParam
	if err := context.ShouldBindQuery(&params); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(params)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errValidation.ErrValidationResponse(err)

		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Err:     err,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     context,
		})
		return
	}

	result, err := f.service.GetFieldSchedule().GetAllWithPagination(context, &params)
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

func (f *FieldScheduleController) GetAllByFieldIDAndDate(context *gin.Context) {
	var params dto.FieldScheduleByFieldIDAndDateRequestParam
	if err := context.ShouldBindQuery(&params); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(params)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errValidation.ErrValidationResponse(err)

		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Err:     err,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     context,
		})
		return
	}

	result, err := f.service.GetFieldSchedule().GetAllByFieldIDAndDate(context, context.Param("uuid"), params.Date)
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

func (f *FieldScheduleController) GetByUUID(context *gin.Context) {
	result, err := f.service.GetFieldSchedule().GetByUUID(context, context.Param("id"))
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

func (f *FieldScheduleController) Create(context *gin.Context) {
	var params dto.FieldScheduleRequest
	err := context.ShouldBindJSON(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Err:     err,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     context,
		})
		return
	}

	err = f.service.GetFieldSchedule().Create(context, &params)
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
		Gin:  context,
	})
}

func (f *FieldScheduleController) Update(context *gin.Context) {
	var params dto.UpdateFieldScheduleRequest
	err := context.ShouldBindJSON(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Err:     err,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     context,
		})
		return
	}

	result, err := f.service.GetFieldSchedule().Update(context, context.Param("uuid"), &params)
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

func (f *FieldScheduleController) UpdateStatus(context *gin.Context) {
	var request dto.UpdateStatusFieldScheduleRequest
	err := context.ShouldBindJSON(&request)
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
			Code:    http.StatusUnprocessableEntity,
			Err:     err,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     context,
		})
		return
	}

	err = f.service.GetFieldSchedule().UpdateStatus(context, &request)
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
		Gin:  context,
	})
}

func (f *FieldScheduleController) Delete(context *gin.Context) {
	err := f.service.GetFieldSchedule().Delete(context, context.Param("uuid"))
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
		Gin:  context,
	})
}

func (f *FieldScheduleController) GenerateScheduleForOneMonth(context *gin.Context) {
	var params dto.GenerateFieldScheduleForOneMonthRequest
	err := context.ShouldBindJSON(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Err:     err,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     context,
		})
		return
	}

	err = f.service.GetFieldSchedule().GenerateScheduleForOneMonth(context, &params)
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
		Gin:  context,
	})
}
