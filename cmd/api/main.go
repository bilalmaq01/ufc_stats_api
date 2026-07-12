package main

import (
	"log"
	"net/http"
	"ufc_stats_api/internal/config"
	"ufc_stats_api/internal/handlers"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/fighters", handlers.GetAllFighters(cfg.DatabaseURL))

	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))

}
