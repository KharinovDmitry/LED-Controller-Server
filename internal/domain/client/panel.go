package client

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"net/url"
)

type Panel interface {
	SendTask(ctx context.Context, host *url.URL, task entity.PanelTask) error
	GetDisplay(ctx context.Context, host *url.URL) (entity.PanelDisplay, error)
}
