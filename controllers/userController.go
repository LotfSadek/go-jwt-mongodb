package controllers

import (
	helper "golang-jwt/helpers"
	"golang-jwt/middleware"
	"golang-jwt/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// put all functions in interface?
// struct controller?
// then struct functions that implement certain interface  functions and return pointer/interface?

type _userHandler struct {
	uc usecases.Base
}

func (c *_userHandler) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := helper.CheckUserType(ctx, "ADMIN"); err != nil {
			return
		}
		result, err := c.uc.GetUsers()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"result": result})
	}
}
func (c *_userHandler) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := helper.CheckUserType(ctx, "ADMIN"); err != nil {
			return
		}
		userId := ctx.Param("user_id")
		if err := helper.MatchUserTypeToUid(ctx, userId); err != nil {
			return
		}
		result, err := c.uc.GetUser(userId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"result": result})
	}
}
func (c *_userHandler) SetupRoutes(rg *gin.RouterGroup) {
	rg.Use(middleware.Authenticate())
	rg.GET("/users", c.GetUsers())
	rg.GET("/users/:user_id", c.GetUser())
}
func NewUserHandler(uc usecases.Base) Controller {
	return &_userHandler{
		uc,
	}
}
