package storage

import (
	"context"
	"path"
	"ufc_stats_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetFighterIDByURL(pool *pgxpool.Pool, url string) (int, error) {
	ctx := context.Background()
	var id int
	hash := path.Base(url)
	err := pool.QueryRow(ctx, "SELECT id FROM fighters WHERE url like '%/' || $1", hash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func GetEventIDByURL(pool *pgxpool.Pool, url string) (int, error) {
	ctx := context.Background()
	var id int
	hash := path.Base(url)
	err := pool.QueryRow(ctx, "SELECT id FROM events WHERE url like '%/' || $1", hash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func InsertFight(pool *pgxpool.Pool, f *models.Fight) error {
	ctx := context.Background()
	_, err := pool.Exec(ctx, "INSERT INTO fights(event_id, fighter1_id, fighter2_id, winner_id, weight_class, method, round, time, time_format, referee,details, is_title, url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) ON CONFLICT (url) DO UPDATE SET event_id = EXCLUDED.event_id, fighter1_id = EXCLUDED.fighter1_id, fighter2_id = EXCLUDED.fighter2_id, winner_id = EXCLUDED.winner_id, weight_class = EXCLUDED.weight_class, method = EXCLUDED.method, round = EXCLUDED.round, time = EXCLUDED.time, time_format = EXCLUDED.time_format, referee = EXCLUDED.referee, details = EXCLUDED.details, is_title = EXCLUDED.is_title", f.EventID, f.Fighter1ID, f.Fighter2ID, f.WinnerID, f.WeightClass, f.Method, f.Round, f.Time, f.TimeFormat, f.Referee, f.Details, f.IsTitle, f.URL)
	if err != nil {
		return err
	}
	return nil
}

func GetFightsByEventID(pool *pgxpool.Pool, eventID int) ([]models.Fight, error) {
	ctx := context.Background()
	rows, err := pool.Query(ctx, "SELECT id, event_id, fighter1_id, fighter2_id, winner_id, weight_class, method, round, time, time_format, referee,details, is_title, url FROM fights WHERE event_id = $1", eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fights := []models.Fight{}
	for rows.Next() {
		var f models.Fight
		err := rows.Scan(&f.ID, &f.EventID, &f.Fighter1ID, &f.Fighter2ID, &f.WinnerID, &f.WeightClass, &f.Method, &f.Round, &f.Time, &f.TimeFormat, &f.Referee, &f.Details, &f.IsTitle, &f.URL)
		if err != nil {
			return nil, err
		}
		fights = append(fights, f)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return fights, nil
}

func GetFightsByFighterID(pool *pgxpool.Pool, fighterID int) ([]models.Fight, error) {
	ctx := context.Background()
	rows, err := pool.Query(ctx, "SELECT id, event_id, fighter1_id, fighter2_id, winner_id, weight_class, method, round, time, time_format, referee,details, is_title, url FROM fights WHERE fighter1_id = $1 OR fighter2_id = $1", fighterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fights := []models.Fight{}
	for rows.Next() {
		var f models.Fight
		err := rows.Scan(&f.ID, &f.EventID, &f.Fighter1ID, &f.Fighter2ID, &f.WinnerID, &f.WeightClass, &f.Method, &f.Round, &f.Time, &f.TimeFormat, &f.Referee, &f.Details, &f.IsTitle, &f.URL)
		if err != nil {
			return nil, err
		}
		fights = append(fights, f)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return fights, nil
}
