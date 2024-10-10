package panel

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) AddPanel(ctx context.Context, panel entity.Panel) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetPanelByUUID(ctx context.Context, uuid uuid.UUID) (entity.Panel, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetPanelByMac(ctx context.Context, mac string) (entity.Panel, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetPanelByKey(ctx context.Context, key string) (entity.Panel, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetPanelsByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]entity.Panel, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) UpdatePanel(ctx context.Context, panel entity.Panel) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) DeletePanel(ctx context.Context, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
