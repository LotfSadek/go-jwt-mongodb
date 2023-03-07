package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	SetupRoutes(rg *gin.RouterGroup)
}
