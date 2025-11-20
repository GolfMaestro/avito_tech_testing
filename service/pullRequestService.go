package service

import (
	"avito_tech_testing/models"
	"avito_tech_testing/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreatePullRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newPullRequest models.PullRequest

	err := json.NewDecoder(r.Body).Decode(&newPullRequest)
	if err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(newPullRequest.PullRequestID)
	fmt.Println(newPullRequest.PullRequestName)
	fmt.Println(newPullRequest.AuthorID)

	repository.AddPullRequestInDB(newPullRequest)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func MergeRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updates struct {
		PullRequestID string `json:"pull_request_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}

	repository.MergeRequestInDB(updates.PullRequestID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
