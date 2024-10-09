package panel

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

func (r Repository) AddPanel(ctx context.Context, user entity.Panel) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetPanelByUUID(ctx context.Context, uuid uuid.UUID) (entity.Panel, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetPanelByMac(ctx context.Context, mac string) (entity.Panel, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) UpdatePanel(ctx context.Context, user entity.Panel) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) DeletePanel(ctx context.Context, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) RegisterPanel(ctx context.Context, panel entity.Panel) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetPanelsByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]entity.Panel, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) SendTaskToPanel(ctx context.Context, panelUUID uuid.UUID, task entity.PanelTask) error {
	//TODO implement me
	panic("implement me")
}
