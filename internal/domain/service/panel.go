package service

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"github.com/google/uuid"
)

type Panel interface {
	RegisterPanel(ctx context.Context, panel entity.Panel) error
	GetPanelsByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]entity.Panel, error)
	SendTaskToPanel(ctx context.Context, panelUUID uuid.UUID, task entity.PanelTask) error
}
