package user

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

func (r *Repository) AddUser(ctx context.Context, user entity.User) error {
	sql := `INSERT INTO users (login, password, role) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, sql, user.Login, user.Password, user.Role)
	if err != nil {
		return fmt.Errorf("[ User Repository ] %w", err)
	}

	return nil
}

func (r *Repository) GetUserByUUID(ctx context.Context, uuid uuid.UUID) (entity.User, error) {
	sql := `SELECT id, login, password, role FROM users WHERE id = $1`
	var user User
	err := r.db.QueryRow(ctx, sql, uuid).Scan(
		&user.UUID,
		&user.Login,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return entity.User{}, fmt.Errorf("[ User Repository ] get user by uuid %w", err)
	}

	return entity.User(user), nil
}

func (r *Repository) GetUserByLogin(ctx context.Context, login string) (entity.User, error) {
	sql := `SELECT id, login, password, role FROM users WHERE login = $1`
	var user User
	err := r.db.QueryRow(ctx, sql, login).Scan(
		&user.UUID,
		&user.Login,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return entity.User{}, fmt.Errorf("[ User Repository ] get user by uuid %w", err)
	}

	return entity.User(user), nil
}

func (r *Repository) UpdateUser(ctx context.Context, user entity.User) error {
	sql := `UPDATE users SET login = $1, password = $2, role = $3 WHERE id = $4`
	_, err := r.db.Exec(ctx, sql, user.Login, user.Password, user.Role, user.UUID)
	if err != nil {
		return fmt.Errorf("[ User Repository ] update user: %w", err)
	}

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, uuid uuid.UUID) error {
	sql := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(ctx, sql, uuid)
	if err != nil {
		return fmt.Errorf("[ User Repository ] delete user: %w", err)
	}

	return nil
}
