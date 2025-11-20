package service

import (
	"avito_tech_testing/repository"
	"encoding/json"
	"net/http"
)

func UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updates struct {
		UserID   string `json:"user_id"`
		IsActive bool   `json:"is_active"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}

	repository.UpdateUserStatus(updates.UserID, updates.IsActive)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
