package config

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type AuthService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authService struct {
}

func NewAuthService() *authService {
	return &authService{}
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func (s *authService) GenerateToken(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return SECRET_KEY, nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
