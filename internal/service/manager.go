package service

import (
	"DynamicLED/config"
	"DynamicLED/internal/domain/service"
	"DynamicLED/internal/repository"
	"DynamicLED/internal/repository/panel"
	"DynamicLED/internal/service/auth"
)

type Manager struct {
	service.Auth
	service.Panel
}

func New(cfg *config.Config, repositories *repository.Manager) *Manager {
	return &Manager{
		Auth:  auth.New(repositories.User, cfg.Auth.JwtSecret, cfg.Auth.Salt, cfg.Auth.RefreshExpireTime, cfg.Auth.AccessExpireTime),
		Panel: panel.New(),
	}
}

func (m *Manager) Init() error {
	return nil
}
