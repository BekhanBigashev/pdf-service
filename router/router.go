package router

import (
	"pdf-service/handlers"
	"pdf-service/router/middlewares"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// Middleware для отлова паник
	router.Use(middlewares.ApiErrorMiddleware())
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	router.GET("/sum", func(c *gin.Context) {
		first, _ := strconv.Atoi(c.Query("first"))
		second, _ := strconv.Atoi(c.Query("second"))

		c.JSON(200, gin.H{
			"result": second + first,
		})
	})

	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	auth.GET("/profile", handlers.Profile)
	auth.POST("/merge", handlers.HandleMergePDF)
	auth.POST("/watermark", handlers.HandleWatermark)

	return router
}
