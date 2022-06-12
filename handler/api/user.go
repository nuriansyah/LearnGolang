package api

import (
	"basicDB/config"
	"basicDB/entity"
	"basicDB/helper"
	"basicDB/models"
	"basicDB/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	service     services.UserInterface
	authService config.AuthService
}

func UserHandlerInit(service services.UserInterface, authService config.AuthService) *userHandler {
	return &userHandler{service, authService}
}

/*
ROUTE: API REGISTER
*/

func (h *userHandler) Register(c *gin.Context) {
	//GET REQUEST REGISTER
	var request entity.RegisterRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errMessage := gin.H{"errors": helper.ErrValidationHandler(err)}
		errResponse := helper.ResponseHandler("Register Validation Failed", http.StatusUnprocessableEntity, "failed", errMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	//SAVE USER DB
	user, err := h.service.Register(request)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("Register Failed", http.StatusBadRequest, "failed", errMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//GENERATE TOKEN
	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GenerateToken Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := models.RegisterData(user, token)
	res := helper.ResponseHandler("Register Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/*
ROUTE: API LOGIN
*/
func (h *userHandler) Login(c *gin.Context) {
	//GET REQUEST LOGIN
	var request entity.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrValidationHandler(err)}
		errResponse := helper.ResponseHandler("Login Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	//GET USER LOGGED
	userLogged, err := h.service.Login(request)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("Login Failed", http.StatusBadRequest, "failed", errMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//GENERATE TOKEN
	token, err := h.authService.GenerateToken(userLogged.ID)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GenerateToken Failed", http.StatusBadRequest, "failed", errMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := models.LoginData(userLogged, token)
	res := helper.ResponseHandler("Login Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
