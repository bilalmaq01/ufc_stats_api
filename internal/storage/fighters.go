package storage

import (
	"context"
	"ufc_stats_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllFighters(pool *pgxpool.Pool) ([]models.Fighter, error) {
	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, "SELECT id, name, nickname, height , weight_class, reach_in, wins, losses, draws FROM fighters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var fighters []models.Fighter
	for rows.Next() {
		var f models.Fighter
		err := rows.Scan(&f.ID, &f.Name, &f.Nickname, &f.Height,
			&f.WeightClass, &f.ReachIn, &f.Wins, &f.Losses, &f.Draws)
		if err != nil {
			return nil, err
		}
		fighters = append(fighters, f)
	}
	return fighters, nil
}
func InsertFighter(f models.Fighter, pool *pgxpool.Pool) error {
	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "INSERT INTO fighters (name, nickname, height, weight_class, reach_in, wins, losses, draws) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) ON CONFLICT (name) DO UPDATE SET wins = EXCLUDED.wins, losses = EXCLUDED.losses, draws = EXCLUDED.draws, nickname = EXCLUDED.nickname, height = EXCLUDED.height, weight_class = EXCLUDED.weight_class, reach_in = EXCLUDED.reach_in", f.Name, f.Nickname, f.Height, f.WeightClass, f.ReachIn, f.Wins, f.Losses, f.Draws)
	if err != nil {
		return err
	}
	return nil
}
