package repository

import "context"

type Token interface {
	AddRefresh(ctx context.Context, login string, refresh string) error
	GetRefresh(ctx context.Context, login string) (string, error)
	DeleteRefresh(ctx context.Context, login string) error
}
