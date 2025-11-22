package tests

import (
	"avito_tech_testing/models"
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

	req := httptest.NewRequest(http.MethodGet, "/team/get?team_name=t1", nil)
	w := httptest.NewRecorder()

	service.GetTeamMembers(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("waiting 200, get:  %d", w.Code)
	}

	var team models.Team

	if err := json.NewDecoder(w.Body).Decode(&team); err != nil {
		t.Fatal("Problem with decode json:", err)
	}

	if len(team.Members) != 0 {
		t.Fatalf("waiting 0, get: %d", len(team.Members))
	}

}

func TestGetTeamNotExist(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE")

	req := httptest.NewRequest(http.MethodGet, "/team/get?team_name=t1", nil)
	w := httptest.NewRecorder()

	service.GetTeamMembers(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("waiting 404, get:  %d", w.Code)
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

	if w.Code != http.StatusCreated {
		t.Fatalf("waiting 201, get:  %d", w.Code)
	}

	var count int
	repository.Pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM teams").Scan(&count)
	if count != 1 {
		t.Fatal("Team not saved in db")
	}

}
