package tests

import (
	"avito_tech_testing/dto"
	"avito_tech_testing/repository"
	"avito_tech_testing/service"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTeam(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE")
	repository.Pool.Exec(context.Background(), "INSERT INTO teams(team_name) VALUES ('t1')")

	req := httptest.NewRequest(http.MethodGet, "/team/get", nil)
	w := httptest.NewRecorder()

	service.GetTeamMembers(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("waiting 200, get:  %d", w.Code)
	}

	var teamMembers []dto.TeamMember

	if err := json.NewDecoder(w.Body).Decode(&teamMembers); err != nil {
		t.Fatal("Problem with decode json:", err)
	}

	if len(teamMembers) != 0 {
		t.Fatalf("waiting 0, get: %d", len(teamMembers))
	}

}
