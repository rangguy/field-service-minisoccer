package services

import (
	"field-service/repositories"
	services "field-service/services/user"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
}

type IServiceRegistry interface {
	GetUser() services.IUserService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry) IServiceRegistry {
	return &Registry{
		repository: repository,
	}
}

func (r *Registry) GetUser() services.IUserService {
	return services.NewUserService(r.repository)
}
