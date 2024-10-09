package handler

import (
	"DynamicLED/internal/repository"
	"DynamicLED/internal/service"
)

type Controller struct {
	Service    *service.Manager
	Repository *repository.Manager
}

func New(services *service.Manager, repositories *repository.Manager) *Controller {
	return &Controller{
		Service:    services,
		Repository: repositories,
	}
}
