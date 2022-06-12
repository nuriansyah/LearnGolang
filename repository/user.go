package repository

import (
	"basicDB/entity"
	"gorm.io/gorm"
)

type UserInterface interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	FindOneByEmail(email string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(model entity.User) (entity.User, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *userRepository) Update(model entity.User) (entity.User, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
func (r *userRepository) FindOneByEmail(email string) (entity.User, error) {
	var model entity.User

	err := r.db.Where("email =?", email).Find(&model).Error
	if err != nil {
		return model, err
	}

	return model, nil
}
