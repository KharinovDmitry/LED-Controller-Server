package token

import (
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

func (r Repository) AddRefresh(ctx context.Context, login string, refresh string) error {
	if err := r.client.Set(login, refresh, 0).Err(); err != nil {
		return fmt.Errorf("[ Token Repoistory ] add refresh err: %w", err)
	}

	return nil
}

func (r Repository) GetRefresh(ctx context.Context, login string) (string, error) {
	var token string
	err := r.client.Get(login).Scan(&token)
	if err != nil {
		return "", fmt.Errorf("[ Token Repoistory ] get refresh err: %w", err)
	}

	return token, nil
}

func (r Repository) DeleteRefresh(ctx context.Context, login string) error {
	if err := r.client.Del(login).Err(); err != nil {
		return fmt.Errorf("[ Token Repoistory ] delete refresh err: %w", err)
	}

	return nil
}
