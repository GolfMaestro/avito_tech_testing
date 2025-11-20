package service

import (
	"avito_tech_testing/models"
	"avito_tech_testing/repository"
	"encoding/json"
	"net/http"
)

func GetTeamMembers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	team_name := r.URL.Query().Get("team_name")

	var team_members []models.TeamMember

	team_members = repository.GetTeamMembersFromDB(team_name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(team_members)

}
