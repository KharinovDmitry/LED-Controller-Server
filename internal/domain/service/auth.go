package service

import (
	"DynamicLED/internal/domain/entity"
	"context"
)

type Auth interface {
	Register(ctx context.Context, login, password string) error
	Login(ctx context.Context, login, password string) (access, refresh string, err error)
	Refresh(ctx context.Context, oldAccess, oldRefresh string) (access, refresh string, err error)
	ParseClaims(ctx context.Context, token string) (entity.Claims, error)
}
