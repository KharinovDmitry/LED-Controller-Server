package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type AuthConfig struct {
	JwtSecret         string        `yaml:"jwtSecret"`
	Salt              string        `yaml:"salt"`
	RefreshExpireTime time.Duration `yaml:"refreshExpireTime"`
	AccessExpireTime  time.Duration `yaml:"accessExpireTime"`
}

type Config struct {
	Port string     `yaml:"port"`
	Auth AuthConfig `yaml:"auth"`
}

func ReadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("[ Config ] open file err: %w", err)
	}

	var cfg Config
	err = yaml.NewEncoder(file).Encode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("[ Config ] encode file err: %w", err)
	}

	return &cfg, nil
}
