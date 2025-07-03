package repositories

import (
	repositories "field-service/repositories/user"
	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetUser() repositories.IUserRepository
}

func NewRepositoryRegistry(db *gorm.DB) *Registry {
	return &Registry{db}
}

func (r *Registry) GetUser() repositories.IUserRepository {
	return repositories.NewUserRepository(r.db)
}
