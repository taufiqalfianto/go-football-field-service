package services

import (
	"user-service/repositories"
	services "user-service/services/user"
)

type Registry struct {
	repositories repositories.IRepositoryRegistry
}

type IServiceRegistry interface {
	GetUser() services.IUserService
}

func NewServiceRegistry(repositories repositories.IRepositoryRegistry) IServiceRegistry {
	return &Registry{
		repositories: repositories,
	}
}

func (r *Registry) GetUser() services.IUserService {
	return services.NewUserService(r.repositories)
}
