package router

import (
	"golang-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func StartGinRouter(port string, ctrl ...controllers.Controller) {
	router := gin.New()
	router.Use(gin.Logger())
	rg := router.Group("/api")
	for i := range ctrl {
		ctrl[i].SetupRoutes(rg)
	}
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})
	router.Run(":" + port)
}
