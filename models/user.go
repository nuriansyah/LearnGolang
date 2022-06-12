package models

import "basicDB/entity"

func RegisterData(user entity.User, token string) entity.ResgisterResponse {
	return entity.ResgisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
}

func LoginData(user entity.User, token string) entity.LoginResponse {
	return entity.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
}
