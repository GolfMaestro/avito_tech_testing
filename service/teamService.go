package service

import (
	"avito_tech_testing/models"
	"avito_tech_testing/repository"
	"avito_tech_testing/utility"
	"encoding/json"
	"net/http"
)

func GetTeamMembers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	team_name := r.URL.Query().Get("team_name")

	team_members, err := repository.GetTeamMembersFromDB(team_name)

	if err != nil {
		utility.Err(w, http.StatusNotFound, "NOT_FOUND", "resource not found")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(team_members)
	}

}

func CreateNewTeam(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTeam models.Team

	err := json.NewDecoder(r.Body).Decode(&newTeam)
	if err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}

	repository.AddNewTeamToDB(newTeam)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
