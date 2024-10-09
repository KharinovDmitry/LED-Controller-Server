package http

import (
	"DynamicLED/internal/http/handler"
	"DynamicLED/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(controller *handler.Controller) *gin.Engine {
	router := gin.New()

	router.GET("/health", controller.Health)

	api := router.Group("/api")
	api.Use(middleware.Auth(controller.Service.Auth))

	auth := api.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		auth.POST("/refresh", controller.Refresh)
	}

	user := api.Group("/user")
	{
		user.PUT("/update")
	}

	panel := api.Group("/panel")
	{
		panel.POST("/register")
		panel.POST("/send")
		panel.GET("/:uuid")
		panel.GET("/user/:uuid")
	}

	return router
}
