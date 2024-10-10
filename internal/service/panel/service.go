package panel

import (
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
	"strconv"
)

type Service struct {
	panel repository.Panel
}

func New(panel repository.Panel) *Service {
	return &Service{panel: panel}
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
	//TODO implement me
	panic("implement me")
}

func (s Service) GetPanelsByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]entity.Panel, error) {
	panels, err := s.panel.GetPanelsByUserUUID(ctx, userUUID)
	if err != nil {
		return nil, fmt.Errorf("[ Panel Service ] get panels by user uuid: %w", err)
	}

	return panels, nil
}

func (s Service) GetPanelByMac(ctx context.Context, mac string) (entity.Panel, error) {
	panel, err := s.panel.GetPanelByMac(ctx, mac)
	if err != nil {
		return entity.Panel{}, fmt.Errorf("[ Panel Service ] get panels by mac: %w", err)
	}

	return panel, nil
}

func (s Service) GetPanelByUUID(ctx context.Context, uuid uuid.UUID) (entity.Panel, error) {
	panel, err := s.panel.GetPanelByUUID(ctx, uuid)
	if err != nil {
		return entity.Panel{}, fmt.Errorf("[ Panel Service ] get panels by mac: %w", err)
	}

	return panel, nil
}

func (s Service) GetPanelDisplayByUUID(ctx context.Context, panelUUID uuid.UUID) (entity.PanelDisplay, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) SyncPanelDisplay(ctx context.Context, panelUUID uuid.UUID) (entity.PanelDisplay, error) {
	//TODO implement me
	panic("implement me")
}

func getKeyByMac(mac string) string {
	hash := sha256.New().Sum([]byte(mac))
	key := int(binary.BigEndian.Uint32(hash)) % constant.PanelKeyLengthMask
	return strconv.Itoa(key)
}
