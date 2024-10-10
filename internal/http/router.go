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
		panel.POST("/create", controller.CreatePanel)
		panel.POST("/register", controller.RegisterPanel)
		panel.POST("/send/:uuid", controller.SendTaskToPanel).Use(middleware.CheckPanelOwningByUUID(controller.Repository.Panel))

		panel.GET("/mac/:mac", controller.GetPanelByMAC).Use(middleware.CheckPanelOwningByMac(controller.Repository.Panel))
		panel.GET("/uuid/:uuid", controller.GetPanelByUUID).Use(middleware.CheckPanelOwningByUUID(controller.Repository.Panel))
		panel.GET("/my", controller.GetPanelByUserUUID)

		panel.GET("/display/:uuid", controller.GetDisplay).Use(middleware.CheckPanelOwningByUUID(controller.Repository.Panel))
		panel.PUT("/display/:uuid", controller.SyncDisplay).Use(middleware.CheckPanelOwningByUUID(controller.Repository.Panel))
	}

	return router
}
