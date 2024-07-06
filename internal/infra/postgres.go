package infra

import (
	"AuthService/internal/config"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

func InitPostgres(cfg *config.Config) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port,
		cfg.Postgres.Username, cfg.Postgres.Password,
		cfg.Postgres.DBName, cfg.Postgres.SSLMode,
	)
	db, err := sqlx.ConnectContext(ctx, "postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
