package services

import (
	"field-service/repositories"
	fieldService "field-service/services/field"
	fieldScheduleService "field-service/services/field_schedule"
	timeServices "field-service/services/time"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
}

type IServiceRegistry interface {
	GetField() fieldService.IFieldService
	GetFieldSchedule() fieldScheduleService.IFieldScheduleService
	GetTime() timeServices.ITimeService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry) IServiceRegistry {
	return &Registry{
		repository: repository,
	}
}

func (r *Registry) GetField() fieldService.IFieldService {
	return fieldService.NewFieldService(r.repository)
}

func (r *Registry) GetFieldSchedule() fieldScheduleService.IFieldScheduleService {
	return fieldScheduleService.NewFieldScheduleService(r.repository)
}

func (r *Registry) GetTime() timeServices.ITimeService {
	return timeServices.NewTimeService(r.repository)
}
