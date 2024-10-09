package repository

import (
	"DynamicLED/config"
	"DynamicLED/internal/domain/repository"
)

type Manager struct {
	repository.User
	repository.Panel
}

func New(*config.Config) *Manager {
	return &Manager{}
}

func (m *Manager) Init() error {
	return nil
}
