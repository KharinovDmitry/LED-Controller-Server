package display

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

type Repository struct {
	client *redis.Client
}

func New(client *redis.Client) *Repository {
	return &Repository{client: client}
}

func (r *Repository) SaveDisplay(ctx context.Context, mac string, display entity.PanelDisplay) error {
	var displayJSON Display
	displayJSON.FromEntity(display)

	if err := r.client.Set(mac, displayJSON, 0).Err(); err != nil {
		return fmt.Errorf("[ Display Repository ] save display err: %w", err)
	}

	return nil
}

func (r *Repository) GetDisplay(ctx context.Context, mac string) (entity.PanelDisplay, error) {
	var displayJSON Display
	if err := r.client.Get(mac).Scan(&displayJSON); err != nil {
		return entity.PanelDisplay{}, fmt.Errorf("[ Display Repository ] get display err: %w", err)
	}

	return displayJSON.ToEntity(), nil
}
