package panel

import (
	"DynamicLED/internal/domain/client"
	"DynamicLED/internal/domain/constant"
	"DynamicLED/internal/domain/entity"
	"DynamicLED/internal/domain/repository"
	"DynamicLED/internal/domain/service"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"strconv"
)

type Service struct {
	panel        repository.Panel
	displayCache repository.Display
	client       client.Panel
}

func New(panel repository.Panel, display repository.Display, client client.Panel) *Service {
	return &Service{
		panel:        panel,
		client:       client,
		displayCache: display,
	}
}

func (s Service) CreatePanel(ctx context.Context, rev int, mac, host string) error {
	err := s.panel.AddPanel(ctx, entity.Panel{
		Rev:  rev,
		Mac:  mac,
		Key:  getKeyByMac(mac),
		Host: host,
	})
	if err != nil {
		return fmt.Errorf("[ Panel Service ] create panel: %w", err)
	}

	return nil
}

func (s Service) RegisterPanel(ctx context.Context, key string, userUUID uuid.UUID) error {
	panel, err := s.panel.GetPanelByKey(ctx, key)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return service.ErrPanelNotFound
		}

		return fmt.Errorf("[ Panel Service ] register panel: %w", err)
	}

	panel.Owner = userUUID
	if err := s.panel.UpdatePanel(ctx, panel); err != nil {
		return fmt.Errorf("[ Panel Service ] update panel: %w", err)
	}

	return nil
}

func (s Service) SendTaskToPanel(ctx context.Context, panelUUID uuid.UUID, task entity.PanelTask) error {
	panel, err := s.panel.GetPanelByUUID(ctx, panelUUID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return service.ErrPanelNotFound
		}
		return fmt.Errorf("[ Panel Service ] send task to panel: %w", err)
	}

	if panel.Host == "" {
		return service.ErrPanelNotRegistered
	}

	host, err := url.Parse(panel.Host)
	if err != nil {
		return fmt.Errorf("[ Panel Service ] send task to panel: %w", err)
	}

	if err := s.client.SendTask(ctx, host, task); err != nil {
		return fmt.Errorf("[ Panel Service ] send task to panel: %w", err)
	}

	// обновляем кэш
	display, err := s.displayCache.GetDisplay(ctx, panel.Mac)
	if err := s.client.SendTask(ctx, host, task); err != nil {
		return fmt.Errorf("[ Panel Service ] send task to panel: %w: %s", service.ErrCacheUpdate, err.Error())
	}

	display.Pixels[task.Position.X*display.Width+task.Position.Y] = task.Color
	if err := s.displayCache.SaveDisplay(ctx, panel.Mac, display); err != nil {
		return fmt.Errorf("[ Panel Service ] send task to panel: %w: %s", service.ErrCacheUpdate, err.Error())
	}

	return nil
}

func (s Service) GetPanelsByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]entity.Panel, error) {
	panels, err := s.panel.GetPanelsByUserUUID(ctx, userUUID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return make([]entity.Panel, 0), nil
		}
		return nil, fmt.Errorf("[ Panel Service ] get panels by user uuid: %w", err)
	}

	return panels, nil
}

func (s Service) GetPanelByMac(ctx context.Context, mac string) (entity.Panel, error) {
	panel, err := s.panel.GetPanelByMac(ctx, mac)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return entity.Panel{}, service.ErrPanelNotFound
		}
		return entity.Panel{}, fmt.Errorf("[ Panel Service ] get panels by mac: %w", err)
	}

	return panel, nil
}

func (s Service) GetPanelByUUID(ctx context.Context, uuid uuid.UUID) (entity.Panel, error) {
	panel, err := s.panel.GetPanelByUUID(ctx, uuid)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return entity.Panel{}, service.ErrPanelNotFound
		}
		return entity.Panel{}, fmt.Errorf("[ Panel Service ] get panels by mac: %w", err)
	}

	return panel, nil
}

func (s Service) GetPanelDisplayByUUID(ctx context.Context, panelUUID uuid.UUID) (entity.PanelDisplay, error) {
	panel, err := s.panel.GetPanelByUUID(ctx, panelUUID)
	if err != nil {
		return entity.PanelDisplay{}, fmt.Errorf("[ Panel Service ] get panel display: %w", err)
	}

	display, err := s.displayCache.GetDisplay(ctx, panel.Mac)
	if err != nil {
		// Не смогли достать из кэша, идем к самой панели
		display, err = s.SyncPanelDisplay(ctx, panelUUID)
		if err != nil {
			return entity.PanelDisplay{}, fmt.Errorf("[ Panel Service ] get panel display: %w", err)
		}
	}

	return display, nil
}

func (s Service) SyncPanelDisplay(ctx context.Context, panelUUID uuid.UUID) (entity.PanelDisplay, error) {
	panel, err := s.panel.GetPanelByUUID(ctx, panelUUID)
	if err != nil {
		return entity.PanelDisplay{}, fmt.Errorf("[ Panel Service ] sync panel display: %w", err)
	}

	if panel.Host == "" {
		return entity.PanelDisplay{}, service.ErrPanelNotRegistered
	}

	host, err := url.Parse(panel.Host)
	if err != nil {
		return entity.PanelDisplay{}, fmt.Errorf("[ Panel Service ] sync panel display: %w", err)
	}

	display, err := s.client.GetDisplay(ctx, host)
	if err != nil {
		return entity.PanelDisplay{}, fmt.Errorf("[ Panel Service ] sync panel display: %w", err)
	}

	if err := s.displayCache.SaveDisplay(ctx, panel.Mac, display); err != nil {
		return entity.PanelDisplay{}, fmt.Errorf("[ Panel Service ] sync panel display: %w", err)
	}

	return display, nil
}

func getKeyByMac(mac string) string {
	hash := sha256.New().Sum([]byte(mac))
	key := int(binary.BigEndian.Uint32(hash)) % constant.PanelKeyLengthMask
	return strconv.Itoa(key)
}
