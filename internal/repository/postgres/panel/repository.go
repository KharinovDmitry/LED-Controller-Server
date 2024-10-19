package panel

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"fmt"
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
	sql := `insert into panels(host, mac, rev, owner) values($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(ctx, sql, panel.Host, panel.Mac, panel.Rev, panel.Owner)
	if err != nil {
		return fmt.Errorf("[ Panel Repository ] add panel: %w", err)
	}

	return nil
}

func (r Repository) GetPanelByUUID(ctx context.Context, uuid uuid.UUID) (entity.Panel, error) {
	sql := `select id, host, mac, rev, owner from panels where id=$1`

	var panel Panel
	err := r.db.QueryRow(ctx, sql, uuid).Scan(
		&panel.UUID,
		&panel.Host,
		&panel.Mac,
		&panel.Rev,
		&panel.Owner,
	)
	if err != nil {
		return entity.Panel{}, fmt.Errorf("[ Panel Repository ] get panel by uuid: %w", err)
	}

	return entity.Panel(panel), nil
}

func (r Repository) GetPanelByMac(ctx context.Context, mac string) (entity.Panel, error) {
	sql := `select id, host, mac, rev, owner from panels where mac=$1`

	var panel Panel
	err := r.db.QueryRow(ctx, sql, mac).Scan(
		&panel.UUID,
		&panel.Host,
		&panel.Mac,
		&panel.Rev,
		&panel.Owner,
	)
	if err != nil {
		return entity.Panel{}, fmt.Errorf("[ Panel Repository ] get panel by mac: %w", err)
	}

	return entity.Panel(panel), nil
}

func (r Repository) GetPanelsByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]entity.Panel, error) {
	sql := `select id, host, mac, rev, owner from panels where owner=$1`

	rows, err := r.db.Query(ctx, sql, userUUID)
	if err != nil {
		return nil, err
	}

	var res []entity.Panel
	for rows.Next() {
		var panel Panel
		err = rows.Scan(
			&panel.UUID,
			&panel.Host,
			&panel.Mac,
			&panel.Rev,
			&panel.Owner,
		)
		if err != nil {
			return nil, fmt.Errorf("[ Panel Repository ] get panels by user uuid: %w", err)
		}

		res = append(res, entity.Panel(panel))
	}

	return res, nil
}

func (r Repository) UpdatePanel(ctx context.Context, panel entity.Panel) error {
	sql := `update panels set host = $1, mac = $2, rev = $3, owner = $4 where id = $5`
	_, err := r.db.Exec(ctx, sql,
		panel.Host,
		panel.Mac,
		panel.Rev,
		panel.Owner,
		panel.UUID,
	)
	if err != nil {
		return fmt.Errorf("[ Panel Repository ] update panel: %w", err)
	}

	return nil
}

func (r Repository) DeletePanel(ctx context.Context, uuid uuid.UUID) error {
	sql := `delete from panels where id = $1`
	_, err := r.db.Exec(ctx, sql, uuid)
	if err != nil {
		return fmt.Errorf("[ Panel Repository ] delete panel: %w", err)
	}

}
