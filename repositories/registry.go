package repositories

import (
	repositories "user-service/repositories/user"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetUser() repositories.IUserRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &registry{db: db}
}

func (r *registry) GetUser() repositories.IUserRepository {
	return repositories.NewUserRepository(r.db)
}
