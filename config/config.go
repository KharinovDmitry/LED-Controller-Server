package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Auth     AuthConfig     `yaml:"auth"`
	Redis    RedisConfig    `yaml:"redis"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type AuthConfig struct {
	JwtSecret         string        `yaml:"jwtSecret"`
	Salt              string        `yaml:"salt"`
	RefreshExpireTime time.Duration `yaml:"refreshExpireTime"`
	AccessExpireTime  time.Duration `yaml:"accessExpireTime"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type PostgresConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

func ReadConfig(path string) (*Config, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
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
