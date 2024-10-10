package app

import (
	"DynamicLED/config"
	"DynamicLED/internal/client/panel"
	"DynamicLED/internal/http"
	"DynamicLED/internal/http/handler"
	"DynamicLED/internal/repository"
	"DynamicLED/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type App struct {
	Port       string
	Controller *handler.Controller
	Router     *gin.Engine
}

func New(cfg *config.Config) *App {
	panelClient := panel.New()
	repositories := repository.New(cfg)
	services := service.New(cfg, repositories, panelClient)

	controller := handler.New(services, repositories, panelClient)
	return &App{
		Port:       cfg.Server.Port,
		Controller: controller,
		Router:     http.SetupRouter(controller),
	}
}

func (app *App) Build() error {
	err := app.Controller.Service.Init()
	if err != nil {
		return fmt.Errorf("[ App ] build err: %w", err)
	}

	err = app.Controller.Repository.Init()
	if err != nil {
		return fmt.Errorf("[ App ] build err: %w", err)
	}

	return nil
}

func (app *App) Run() error {
	if err := app.Router.Run(app.Port); err != nil {
		return fmt.Errorf("[ App ] run err: %w", err)
	}

	return nil
}
