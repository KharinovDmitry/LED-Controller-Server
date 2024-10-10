package postgres

import (
	"DynamicLED/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

var ConnStrTemplate = "postgres://%s:%s@%s:%s/%s"

func GetConn(cfg *config.PostgresConfig) (*pgx.Conn, error) {
	return pgx.Connect(context.TODO(), fmt.Sprintf(ConnStrTemplate, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
}
