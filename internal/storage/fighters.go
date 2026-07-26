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
	rows, err := conn.Query(ctx, "SELECT id, name, nickname, height , weight, reach, wins, losses, draws FROM fighters")
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
	_, err = conn.Exec(ctx, "INSERT INTO fighters (name, nickname, height, weight, reach, wins, losses, draws, stance, dob, slpm, str_acc, sapm, str_def, td_avg, td_acc, td_def, sub_avg, url) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19) ON CONFLICT (url) DO UPDATE SET name = EXCLUDED.name, nickname = EXCLUDED.nickname, height = EXCLUDED.height, weight = EXCLUDED.weight, reach = EXCLUDED.reach, wins = EXCLUDED.wins, losses = EXCLUDED.losses, draws = EXCLUDED.draws, stance = EXCLUDED.stance, dob = EXCLUDED.dob, slpm = EXCLUDED.slpm, str_acc = EXCLUDED.str_acc, sapm = EXCLUDED.sapm, str_def = EXCLUDED.str_def, td_avg = EXCLUDED.td_avg, td_acc = EXCLUDED.td_acc, td_def = EXCLUDED.td_def, sub_avg = EXCLUDED.sub_avg", f.Name, f.Nickname, f.Height, f.WeightClass, f.ReachIn, f.Wins, f.Losses, f.Draws, f.Stance, f.DOB, f.SLPM, f.StrAcc, f.SAPM, f.StrDef, f.TdAvg, f.TdAcc, f.TdDef, f.SubAvg, f.URL)
	if err != nil {
		return err
	}
	return nil
}

// Function for getting fighter info by name, returns a slice with all fighters with a matching name
func GetFighterByname(name string, pool *pgxpool.Pool) ([]models.Fighter, error) {
	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return []models.Fighter{}, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, "SELECT id, name, nickname, height , weight, reach, wins, losses, draws, stance, dob, slpm, str_acc, sapm, str_def, td_avg, td_acc, td_def, sub_avg, url FROM fighters WHERE LOWER(name) LIKE LOWER($1)", "%"+name+"%")
	if err != nil {
		return []models.Fighter{}, err
	}

	var fighters []models.Fighter
	//Scans next rows for more fighters
	for rows.Next() {
		var f models.Fighter
		err := rows.Scan(&f.ID, &f.Name, &f.Nickname, &f.Height, &f.WeightClass, &f.ReachIn, &f.Wins, &f.Losses, &f.Draws, &f.Stance, &f.DOB, &f.SLPM, &f.StrAcc, &f.SAPM, &f.StrDef, &f.TdAvg, &f.TdAcc, &f.TdDef, &f.SubAvg, &f.URL)
		if err != nil {
			return nil, err
		}
		fighters = append(fighters, f)
	}
	return fighters, nil
}
