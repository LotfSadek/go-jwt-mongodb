package controllers

import (
	"golang-jwt/models"
	"golang-jwt/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type _authHandler struct {
	uc usecases.Base
}

func (c *_authHandler) SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := c.uc.SignUp(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"result": result})
	}
}
func (c *_authHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := c.uc.Login(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"result": result})
	}
}
func (c *_authHandler) SetupRoutes(rg *gin.RouterGroup) {
	rg.POST("/users/signup", c.SignUp())
	rg.POST("/users/login", c.Login())
}
func NewAuthHandler(uc usecases.Base) Controller {
	return &_authHandler{
		uc,
	}
}
