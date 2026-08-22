package handlers

import (
	"encoding/json"
	"net/http"
	"ufc_stats_api/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllEvents(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := storage.GetAllEvents(pool)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
	}

}
