package repository

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"github.com/google/uuid"
)

type User interface {
	AddUser(ctx context.Context, user entity.User) error
	GetUserByUUID(ctx context.Context, uuid uuid.UUID) (entity.User, error)
	GetUserByLogin(ctx context.Context, login string) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) error
	DeleteUser(ctx context.Context, uuid uuid.UUID) error
}
