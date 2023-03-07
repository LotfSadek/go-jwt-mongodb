package controllers

import (
	helper "golang-jwt/helpers"
	"golang-jwt/middleware"
	"golang-jwt/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// put all functions in interface?
// struct controller?
// then struct functions that implement certain interface  functions and return pointer/interface?

type _userHandler struct {
	uc usecases.Base
}

func (c *_userHandler) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var allusers []primitive.M
		if err := helper.CheckUserType(ctx, "ADMIN"); err != nil {
			return
		}
		if err := ctx.BindJSON(&allusers); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := c.uc.GetUsers(allusers)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"result": result})
	}
}
func (c *_userHandler) GetUser() gin.HandlerFunc {
	return c.uc.GetUser()
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
