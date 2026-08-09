package storage

import (
	"context"
	"ufc_stats_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertEvent(pool *pgxpool.Pool, e *models.Event) error {
	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "INSERT INTO events (event_name, date, location, url) VALUES ($1, $2, $3, $4) ON CONFLICT (url) DO UPDATE SET event_name = EXCLUDED.event_name, date = EXCLUDED.date, location = EXCLUDED.location", e.EventName, e.Date, e.Location, e.URL)
	if err != nil {
		return err
	}
	return nil
}
