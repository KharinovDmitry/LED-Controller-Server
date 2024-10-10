package panel

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"net/http"
	"net/url"
)

type Panel struct {
	client *http.Client
}

func New() *Panel {
	return &Panel{client: http.DefaultClient}
}

func (p *Panel) SendTask(ctx context.Context, host *url.URL, task entity.PanelTask) error {
	return nil
}

func (p *Panel) GetDisplay(ctx context.Context, host *url.URL) (entity.PanelDisplay, error) {
	return entity.PanelDisplay{}, nil
}
