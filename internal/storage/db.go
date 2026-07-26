package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect to database
func Connect(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, databaseURL)
}
