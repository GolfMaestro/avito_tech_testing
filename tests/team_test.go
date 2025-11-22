package tests

import (
	"avito_tech_testing/dto"
	"avito_tech_testing/repository"
	"avito_tech_testing/service"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestAddTeam(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE")

	values := strings.NewReader("{\"team_name\": \"t1\", \"members\": []}")
	req := httptest.NewRequest(http.MethodPost, "/persons", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.CreateNewTeam(w, req)

	// if w.Code != http.StatusCreated {
	// 	t.Fatalf("waiting 201, get:  %d", w.Code)
	// }

	var count int
	repository.Pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM teams").Scan(&count)
	if count != 1 {
		t.Fatal("Team not saved in db")
	}

}
