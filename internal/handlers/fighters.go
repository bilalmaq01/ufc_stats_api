package handlers

import (
	"encoding/json"
	"net/http"
	"ufc_stats_api/internal/storage"
)

func GetAllFighters(w http.ResponseWriter, r *http.Request) {
	fighters, err := storage.GetAllFighters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fighters)
}
