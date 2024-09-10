package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/config"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/usecase"
	"working.com/bank_dash/package/tokens"
)

type UserController struct {
	UserUseCase *usecase.UserUseCase
	Env         *config.Env
}

func NewUserController(env *config.Env, useCase *usecase.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: useCase,
		Env:         env,
	}
}

// handler for handling login process
func (uc *UserController) LoginIN(c *gin.Context) {
	var loginData domain.LoginRequest
	if err := c.BindJSON(&loginData); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Invalid data"})
		return
	}

	user, err := uc.UserUseCase.LoginUser(c, &loginData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	access, err := tokens.CreateAccessToken(user, uc.Env.AccessTokenSecret, uc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	refresh, err := tokens.CreateRefreshToken(user, uc.Env.RefreshTokenSecret, uc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.SetCookie(
		"refresh_token",
		refresh,
		uc.Env.RefreshTokenExpiryHour,
		"",
		"",
		true,
		true,
	)

	response := map[string]interface{}{
		"success": true,
		"message": "logged in",
		"data": map[string]interface{}{
			"access_token":  access,
			"refresh_token": refresh,
			"data":          map[string]interface{}{},
		},
	}
	c.IndentedJSON(http.StatusOK, response)
}

// Signup handler for registering a new user
func (uc *UserController) Signup(c *gin.Context) {
	var user *domain.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	registered, err := uc.UserUseCase.RegisterUser(c, user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	access, err := tokens.CreateAccessToken(user, uc.Env.AccessTokenSecret, uc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	refresh, err := tokens.CreateRefreshToken(user, uc.Env.RefreshTokenSecret, uc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success":       true,
		"message":       "succesfull registration",
		"access_token":  access,
		"refresh_token": refresh,
		"data":          registered,
	}

	c.IndentedJSON(http.StatusOK, response)
}

// method for updating the user profile
func (uc *UserController) UpdateProfile(c *gin.Context) {
	var user domain.UserRequest
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	Id, exist := c.Get("id")
	if !exist {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error of id"})
		return
	}

	id, ok := Id.(string)
	fmt.Print(id)
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}

	updated, err := uc.UserUseCase.UpdateInfo(c, id, &user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	response := map[string]interface{}{
		"success": true,
		"message": "profile is updated",
		"data":    updated,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// method for updating the user preference
func (uc *UserController) UpdatePreference(c *gin.Context) {
	var preference domain.UserPreference
	if err := c.BindJSON(&preference); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	Id, exist := c.Get("id")
	if !exist {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error of id"})
		return
	}

	id, ok := Id.(string)
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}

	updated, err := uc.UserUseCase.UpdatePreference(c, id, &preference)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	response := map[string]interface{}{
		"success": true,
		"message": "profile is updated",
		"data":    updated,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// method to get user by using username
func (uc *UserController) GetByUserName(c *gin.Context) {
	username := c.Param("username")

	if username == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "username name is needed"})
		return
	}

	user, err := uc.UserUseCase.GetByUserName(c, username)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "user information",
		"data":    user,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler for refreshing the token
func (uc *UserController) RefreshToken(c *gin.Context) {
	userData, okay := c.Get("username")
	if !okay {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error of data"})
		return
	}
	username, okay := userData.(string)
	if !okay {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error of data type"})
		return
	}

	refreshToken, err := c.Cookie("refresh-token")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "no existing  data"})
		return
	}

	_, err = tokens.VerifyToken(refreshToken, uc.Env.RefreshTokenSecret)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "login again"})
		return
	}

	user, err := uc.UserUseCase.GetByUserNameForPass(c, username)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	access, err := tokens.CreateAccessToken(user, uc.Env.AccessTokenSecret, uc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "created",
		"data": map[string]interface{}{
			"access_token":  access,
			"refresh_token": refreshToken,
			"data":          map[string]interface{}{},
		},
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler for working with changing password or username
func (uc *UserController) ChangePassword(c *gin.Context) {
	var loginData domain.ChangePassword
	if err := c.BindJSON(&loginData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	userData, okay := c.Get("username")
	if !okay {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	username, okay := userData.(string)
	if !okay {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of data"})
		return
	}

	err := uc.UserUseCase.UpdatePassword(c, username, &loginData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"service": true,
		"message": "updated",
		"data":    map[string]interface{}{},
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler for getting current user information

func (uc *UserController) GetCurrentUser(c *gin.Context) {
	userdata, exist := c.Get("username")
	if userdata == "" || !exist {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	username, okay := userdata.(string)
	if !okay {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Invalid data"})
		return
	}

	user, err := uc.UserUseCase.GetByUserName(c, username)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "user information",
		"data":    user,
	}
	c.IndentedJSON(http.StatusOK, response)

}
