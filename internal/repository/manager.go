package repository

import (
	"DynamicLED/config"
	"DynamicLED/internal/domain/repository"
	"DynamicLED/internal/repository/postgres"
	"DynamicLED/internal/repository/postgres/panel"
	"DynamicLED/internal/repository/postgres/user"
	"DynamicLED/internal/repository/redis"
	"DynamicLED/internal/repository/redis/token"
	"fmt"
)

type Manager struct {
	RedisConfig    *config.RedisConfig
	PostgresConfig *config.PostgresConfig

	repository.User
	repository.Panel
	repository.Token
}

func New(cfg *config.Config) *Manager {
	return &Manager{
		RedisConfig:    &cfg.Redis,
		PostgresConfig: &cfg.Postgres,
	}
}

func (m *Manager) Init() error {
	redisClient := redis.NewClient(m.RedisConfig)
	postgresConn, err := postgres.GetConn(m.PostgresConfig)
	if err != nil {
		return fmt.Errorf("[ Repository Manager ] init err: %w", err)
	}

	m.User = user.New(postgresConn)
	m.Panel = panel.New(postgresConn)
	m.Token = token.New(redisClient)

	return nil
}
