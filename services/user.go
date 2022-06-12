package services

import (
	"basicDB/entity"
	"basicDB/repository"
	"errors"
)

type UserInterface interface {
	Login(form entity.LoginRequest) (entity.User, error)
	Register(form entity.RegisterRequest) (entity.User, error)
}

type userService struct {
	repository repository.UserInterface
}

func NewUserServices(repository repository.UserInterface) *userService {
	return &userService{repository}
}

func Login() {

}
func (s *userService) Register(form entity.RegisterRequest) (entity.User, error) {
	var model entity.User

	model = entity.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	newUser, err := s.repository.Create(model)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
func (s *userService) Login(form entity.LoginRequest) (entity.User, error) {

	model, err := s.repository.FindOneByEmail(form.Email)
	if err != nil {
		return model, err
	}

	if model.ID == 0 {
		return model, errors.New("User not found")
	}
	return model, nil
}
