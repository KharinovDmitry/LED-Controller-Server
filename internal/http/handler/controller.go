package handler

import (
	"DynamicLED/internal/domain/client"
	"DynamicLED/internal/repository"
	"DynamicLED/internal/service"
)

type Controller struct {
	Service    *service.Manager
	Repository *repository.Manager

	PanelClient client.Panel
}

func New(services *service.Manager, repositories *repository.Manager, panelClient client.Panel) *Controller {
	return &Controller{
		Service:     services,
		Repository:  repositories,
		PanelClient: panelClient,
	}
}
