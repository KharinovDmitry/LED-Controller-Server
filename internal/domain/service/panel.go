package service

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"github.com/google/uuid"
)

type Panel interface {
	// CreatePanel панель заводится в бд
	CreatePanel(ctx context.Context, rev int, mac, host string) error

	// RegisterPanel пользователь регистрирует панель на себя по уникальному ключу
	RegisterPanel(ctx context.Context, key string, userUUID uuid.UUID) error

	// SendTaskToPanel отправка задачи на панель
	SendTaskToPanel(ctx context.Context, panelUUID uuid.UUID, task entity.PanelTask) error

	GetPanelsByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]entity.Panel, error)
	GetPanelByMac(ctx context.Context, mac string) (entity.Panel, error)
	GetPanelByUUID(ctx context.Context, uuid uuid.UUID) (entity.Panel, error)

	// GetPanelDisplayByUUID как сейчас выглядит дисплей матрицы, данные НЕ с панели
	GetPanelDisplayByUUID(ctx context.Context, panelUUID uuid.UUID) (entity.PanelDisplay, error)

	// SyncPanelDisplay как сейчас выглядит дисплей матрицы, данные с панели
	SyncPanelDisplay(ctx context.Context, panelUUID uuid.UUID) (entity.PanelDisplay, error)
}
