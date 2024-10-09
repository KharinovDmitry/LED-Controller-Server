package dto

import (
	"DynamicLED/internal/domain/service"
)

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (c *Credentials) Validate() error {
	if c.Login == "" {
		return service.ErrInvalidCredentials
	}
	if c.Password == "" {
		return service.ErrInvalidCredentials
	}
	return nil
}

type TokenPair struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func NewTokenPair(access, refresh string) TokenPair {
	return TokenPair{
		Access:  access,
		Refresh: refresh,
	}
}
