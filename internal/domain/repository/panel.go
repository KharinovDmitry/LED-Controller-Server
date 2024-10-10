package repository

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"github.com/google/uuid"
)

type Panel interface {
	AddPanel(ctx context.Context, panel entity.Panel) error
	GetPanelByUUID(ctx context.Context, uuid uuid.UUID) (entity.Panel, error)
	GetPanelByMac(ctx context.Context, mac string) (entity.Panel, error)
	GetPanelByKey(ctx context.Context, key string) (entity.Panel, error)
	GetPanelsByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]entity.Panel, error)
	UpdatePanel(ctx context.Context, panel entity.Panel) error
	DeletePanel(ctx context.Context, uuid uuid.UUID) error
}
