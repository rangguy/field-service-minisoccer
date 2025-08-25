package controllers

import (
	errValidation "field-service/common/error"
	"field-service/common/response"
	"field-service/domain/dto"
	fieldService "field-service/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type FieldController struct {
	service fieldService.IServiceRegistry
}

type IFieldController interface {
	GetAllWithPagination(*gin.Context)
	GetAllWithoutPagination(*gin.Context)
	GetByUUID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func NewFieldController(service fieldService.IServiceRegistry) IFieldController {
	return &FieldController{service: service}
}

func (f *FieldController) GetAllWithPagination(ctx *gin.Context) {
	var params dto.FieldRequestParam
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

	result, err := f.service.GetField().GetAllWithPagination(ctx, &params)
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

func (f *FieldController) GetAllWithoutPagination(ctx *gin.Context) {
	result, err := f.service.GetField().GetAllWithoutPagination(ctx)
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

func (f *FieldController) GetByUUID(ctx *gin.Context) {
	result, err := f.service.GetField().GetByUUID(ctx, ctx.Param("uuid"))
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

func (f *FieldController) Create(ctx *gin.Context) {
	var request dto.FieldRequest
	if err := ctx.ShouldBindWith(&request, binding.FormMultipart); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
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

	result, err := f.service.GetField().Create(ctx, &request)
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

func (f *FieldController) Update(ctx *gin.Context) {
	var request dto.UpdateFieldRequest
	if err := ctx.ShouldBindWith(&request, binding.FormMultipart); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
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

	result, err := f.service.GetField().Update(ctx, ctx.Param("uuid"), &request)
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

func (f *FieldController) Delete(ctx *gin.Context) {
	err := f.service.GetField().Delete(ctx, ctx.Param("uuid"))
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
