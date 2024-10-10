package repository

import (
	"DynamicLED/internal/domain/entity"
	"context"
)

type Display interface {
	SaveDisplay(ctx context.Context, mac string, display entity.PanelDisplay) error
	GetDisplay(ctx context.Context, mac string) (entity.PanelDisplay, error)
}
