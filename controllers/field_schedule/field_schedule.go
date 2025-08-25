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

func (f *FieldScheduleController) GetAllWithPagination(ctx *gin.Context) {
	var params dto.FieldScheduleRequestParam
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
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
			Gin:     ctx,
		})
		return
	}

	result, err := f.service.GetFieldSchedule().GetAllWithPagination(ctx, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})
}

func (f *FieldScheduleController) GetAllByFieldIDAndDate(ctx *gin.Context) {
	var params dto.FieldScheduleByFieldIDAndDateRequestParam
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
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
			Gin:     ctx,
		})
		return
	}

	result, err := f.service.GetFieldSchedule().GetAllByFieldIDAndDate(ctx, ctx.Param("uuid"), params.Date)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})
}

func (f *FieldScheduleController) GetByUUID(ctx *gin.Context) {
	result, err := f.service.GetFieldSchedule().GetByUUID(ctx, ctx.Param("uuid"))
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})
}

func (f *FieldScheduleController) Create(ctx *gin.Context) {
	var params dto.FieldScheduleRequest
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
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
			Gin:     ctx,
		})
		return
	}

	err = f.service.GetFieldSchedule().Create(ctx, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusCreated,
		Gin:  ctx,
	})
}

func (f *FieldScheduleController) Update(ctx *gin.Context) {
	var params dto.UpdateFieldScheduleRequest
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
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
			Gin:     ctx,
		})
		return
	}

	result, err := f.service.GetFieldSchedule().Update(ctx, ctx.Param("uuid"), &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})
}

func (f *FieldScheduleController) UpdateStatus(ctx *gin.Context) {
	var request dto.UpdateStatusFieldScheduleRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
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
			Gin:     ctx,
		})
		return
	}

	err = f.service.GetFieldSchedule().UpdateStatus(ctx, &request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Gin:  ctx,
	})
}

func (f *FieldScheduleController) Delete(ctx *gin.Context) {
	err := f.service.GetFieldSchedule().Delete(ctx, ctx.Param("uuid"))
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Gin:  ctx,
	})
}

func (f *FieldScheduleController) GenerateScheduleForOneMonth(ctx *gin.Context) {
	var params dto.GenerateFieldScheduleForOneMonthRequest
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
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
			Gin:     ctx,
		})
		return
	}

	err = f.service.GetFieldSchedule().GenerateScheduleForOneMonth(ctx, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusCreated,
		Gin:  ctx,
	})
}
