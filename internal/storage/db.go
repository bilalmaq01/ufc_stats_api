package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://apple@localhost:5432/ufc_stats")
}
