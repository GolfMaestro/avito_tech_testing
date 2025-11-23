package service

import (
	"avito_tech_testing/dto"
	"avito_tech_testing/models"
	"avito_tech_testing/repository"
	"avito_tech_testing/utility"
	"encoding/json"
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

	pullRequest, err := repository.AddPullRequestInDB(newPullRequest)

	if err == nil {
		response := struct {
			PullRequest dto.CreatePullRequestResponse `json:"pr"`
		}{PullRequest: pullRequest}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	} else if err.Error() == "NOT_FOUND" {
		utility.Err(w, http.StatusNotFound, "NOT_FOUND", "resource not found")
	} else if err.Error() == "PR_EXISTS" {
		utility.Err(w, http.StatusConflict, "PR_EXISTS", "PR id already exists")
	}

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

	mergePullRequst, err := repository.MergeRequestInDB(updates.PullRequestID)

	if err == nil {
		response := struct {
			PullRequest dto.MergePullRequestResponse `json:"pr"`
		}{PullRequest: mergePullRequst}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	} else if err.Error() == "NOT_FOUND" {
		utility.Err(w, http.StatusNotFound, "NOT_FOUND", "resource not found")
	} else if err.Error() == "PR_ALREADY_MERGED" {
		utility.Err(w, http.StatusConflict, "PR_ALREADY_MERGED", "PR already merged")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func ReassignRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updates struct {
		PullRequestID string `json:"pull_request_id"`
		OldReviewerID string `json:"old_reviewer_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}

	repository.ReassignRequestInDB(updates.PullRequestID, updates.OldReviewerID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
