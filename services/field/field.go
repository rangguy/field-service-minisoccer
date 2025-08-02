package services

import (
	"context"
	"field-service/common/util"
	"field-service/domain/dto"
	"field-service/repositories"
)

type FieldService struct {
	repository repositories.IRepositoryRegistry
}

type IFieldService interface {
	GetAllWithPagination(context.Context, *dto.FieldRequestParam) (*util.PaginationResult, error)
	GetAllWithoutPagination(context.Context) ([]dto.FieldResponse, error)
	GetByUUID(context.Context, string) (*dto.FieldResponse, error)
	Create(context.Context, *dto.FieldRequest) (*dto.FieldResponse, error)
	Update(context.Context, string, *dto.UpdateFieldRequest) (*dto.FieldResponse, error)
	Delete(context.Context, string) error
}

func NewFieldService(repository repositories.IRepositoryRegistry) *FieldService {
	return &FieldService{repository}
}

func (f *FieldService) GetAllWithPagination(ctx context.Context, param *dto.FieldRequestParam) (*util.PaginationResult, error) {
	fields, total, err := f.repository.GetField().FindAllWithPagination(ctx, param)
	if err != nil {
		return nil, err
	}

	fieldResult := make([]*dto.FieldResponse, 0, len(fields))
	for _, field := range fields {
		fieldResult = append(fieldResult, &dto.FieldResponse{
			UUID:         field.UUID,
			Name:         field.Name,
			PricePerHour: field.PricePerHour,
			Images:       field.Images,
			CreatedAt:    field.CreatedAt,
			UpdatedAt:    field.UpdatedAt,
		})
	}

	pagination := &util.PaginationParam{
		Count: total,
		Page:  param.Page,
		Limit: param.Limit,
		Data:  fieldResult,
	}

	response := util.GeneratePagination(*pagination)
	return response, nil
}

func (f *FieldService) GetAllWithoutPagination(ctx context.Context) ([]dto.FieldResponse, error) {
	fields, err := f.repository.GetField().FindAllWithoutPagination(ctx)
	if err != nil {
		return nil, err
	}

	fieldResult := make([]dto.FieldResponse, 0, len(fields))
	for _, field := range fields {
		fieldResult = append(fieldResult, dto.FieldResponse{
			UUID:         field.UUID,
			Name:         field.Name,
			PricePerHour: field.PricePerHour,
			Images:       field.Images,
		})
	}

	return fieldResult, nil
}

func (f *FieldService) GetByUUID(ctx context.Context, uuid string) (*dto.FieldResponse, error) {
	field, err := f.repository.GetField().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	fieldResult := &dto.FieldResponse{
		UUID:         field.UUID,
		Name:         field.Name,
		PricePerHour: field.PricePerHour,
		Images:       field.Images,
		CreatedAt:    field.CreatedAt,
		UpdatedAt:    field.UpdatedAt,
	}

	return fieldResult, nil
}

func (f *FieldService) Create(ctx context.Context, request *dto.FieldRequest) (*dto.FieldResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FieldService) Update(ctx context.Context, s string, request *dto.UpdateFieldRequest) (*dto.FieldResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FieldService) Delete(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}
