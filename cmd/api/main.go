package main

import (
	"context"
	"log"
	"net/http"
	"ufc_stats_api/internal/config"
	"ufc_stats_api/internal/handlers"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/fighters", handlers.GetAllFighters(pool))

	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))

}
