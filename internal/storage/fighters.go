package storage

import (
	"context"
	"ufc_stats_api/internal/models"
)

func GetAllFighters(databaseURL string) ([]models.Fighter, error) {
	ctx := context.Background()
	conn, err := Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)
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
func InsertFighter(f models.Fighter, databaseURL string) error {
	ctx := context.Background()
	conn, err := Connect(ctx, databaseURL)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	_, err = conn.Exec(ctx, "INSERT INTO fighters (name, nickname, height, weight_class, reach_in, wins, losses, draws) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", f.Name, f.Nickname, f.Height, f.WeightClass, f.ReachIn, f.Wins, f.Losses, f.Draws)
	if err != nil {
		return err
	}
	return nil
}
