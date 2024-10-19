package client

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"net/url"
)

type Panel interface {
	SendTask(ctx context.Context, host *url.URL, task entity.PanelTask) error

	// SendTaskButch НЕ гарантирует порядок отправки задач
	SendTaskButch(ctx context.Context, host *url.URL, tasks []entity.PanelTask) (entity.ButchReport, error)

	GetDisplay(ctx context.Context, host *url.URL) (entity.PanelDisplay, error)
}
