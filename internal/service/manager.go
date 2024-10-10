package service

import (
	"DynamicLED/config"
	"DynamicLED/internal/domain/client"
	"DynamicLED/internal/domain/service"
	"DynamicLED/internal/repository"
	"DynamicLED/internal/service/auth"
	"DynamicLED/internal/service/panel"
)

type Manager struct {
	service.Auth
	service.Panel
}

func New(cfg *config.Config, repositories *repository.Manager, panelClient client.Panel) *Manager {
	return &Manager{
		Auth:  auth.New(repositories.User, repositories.Token, cfg.Auth),
		Panel: panel.New(repositories.Panel, repositories.Display, panelClient),
	}
}

func (m *Manager) Init() error {
	return nil
}
