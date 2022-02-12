package repository

import (
	"github.com/den19980107/go-fiber-gorm-starter/db/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	ORM *gorm.DB
}

type UserRepositoryInterface interface {
	GetByUsername(username string) *entity.User
	Create(user *entity.User) *entity.User
}

// Create a new user repository instance.
func New(orm *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		ORM: orm,
	}
}

func (repo *UserRepository) GetByUsername(username string) *entity.User {
	var user entity.User

	if err := repo.ORM.Where(&entity.User{Username: username}).First(&user).Error; err != nil {
		return nil
	}

	return &user
}

func (repo *UserRepository) Create(user *entity.User) *entity.User {
	repo.ORM.Create(user)

	return user
}
