package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"ufc_stats_api/internal/storage"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllFighters(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fighters, err := storage.GetAllFighters(pool)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fighters)
	}
}
func SearchFightersByName(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "name query parameter is required", http.StatusBadRequest)
			return
		}
		fighter, err := storage.GetFighterByname(name, pool)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				http.Error(w, "fighter not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fighter)
	}
}
