package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"

	"chat_server/internal/client/db"
)

type pgClient struct {
	masterDBC db.DB
}

// New создаёт объект клиента бд
func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

// DB является геттером для мастер бд
func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

// Close закрывает соединение к бд
func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}