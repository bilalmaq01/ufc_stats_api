package storage

import (
	"context"
	"ufc_stats_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertEvent(pool *pgxpool.Pool, e *models.Event) error {
	ctx := context.Background()
	_, err := pool.Exec(ctx, "INSERT INTO events (event_name, date, location, url) VALUES ($1, $2, $3, $4) ON CONFLICT (url) DO UPDATE SET event_name = EXCLUDED.event_name, date = EXCLUDED.date, location = EXCLUDED.location", e.EventName, e.Date, e.Location, e.URL)
	if err != nil {
		return err
	}
	return nil
}
func GetAllEvents(pool *pgxpool.Pool) ([]models.Event, error) {
	ctx := context.Background()
	rows, err := pool.Query(ctx, "SELECT id,event_name, date, location, url FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []models.Event{}
	for rows.Next() {
		var e models.Event
		err := rows.Scan(&e.ID, &e.EventName, &e.Date, &e.Location, &e.URL)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
