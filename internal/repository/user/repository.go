package user

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"github.com/google/uuid"
)

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r Repository) AddUser(ctx context.Context, user entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetUserByUUID(ctx context.Context, uuid uuid.UUID) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetUserByLogin(ctx context.Context, login string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) UpdateUser(ctx context.Context, user entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) DeleteUser(ctx context.Context, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
