package service

import (
	"avito_tech_testing/dto"
	"avito_tech_testing/repository"
	"avito_tech_testing/utility"
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

	user, err := repository.UpdateUserStatus(updates.UserID, updates.IsActive)

	if err != nil {
		utility.Err(w, http.StatusNotFound, "NOT_FOUND", "resource not found")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}

}

func GetUserReviews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	user_id := r.URL.Query().Get("user_id")

	var pullRequests []dto.PullRequestShort

	pullRequests = repository.GetUserReviewsFromDB(user_id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pullRequests)

}
