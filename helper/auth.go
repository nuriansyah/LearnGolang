package helper

import (
	"basicDB/entity"
	"encoding/json"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserLoggedin(c *gin.Context) (entity.User, error) {
	session := sessions.Default(c)
	var user entity.User
	userSession := session.Get(user)
	if userSession == nil {
		c.Redirect(http.StatusNotFound, "/login")
		return entity.User{}, errors.New("User not Loggedin")
	}
	err := json.Unmarshal(userSession.([]byte), &user)
	if err != nil {
		c.Redirect(http.StatusNotFound, "/login")
		return entity.User{}, err
	}
	return user, nil
}
