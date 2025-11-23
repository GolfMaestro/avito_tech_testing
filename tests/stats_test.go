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

func TestGetAllPRStats(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE pull_requests CASCADE")
	repository.Pool.Exec(context.Background(), "INSERT INTO teams(team_name) VALUES ('t1');")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u1', 'bot1', 't1', true);")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u125', 'bot125', 't1', true);")
	repository.Pool.Exec(context.Background(), "INSERT INTO pull_requests(pull_request_id, pull_request_name, author_id, assigned_reviewers) VALUES ('pr1', 'update_1', 'u125', ARRAY['u1']);")
	repository.Pool.Exec(context.Background(), "INSERT INTO pull_requests(pull_request_id, pull_request_name, author_id, assigned_reviewers) VALUES ('pr2', 'update_1', 'u125', ARRAY['u1']);")
	repository.Pool.Exec(context.Background(), "INSERT INTO pull_requests(pull_request_id, pull_request_name, author_id, assigned_reviewers) VALUES ('pr3', 'update_1', 'u125', ARRAY['u1']);")
	repository.Pool.Exec(context.Background(), "UPDATE pull_requests SET status = 'MERGED', merged_at = NOW() WHERE pull_request_id = 'pr1';")

	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	w := httptest.NewRecorder()

	service.GetStats(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("waiting 200, get:  %d", w.Code)
	}

	var stats dto.Stats

	if err := json.NewDecoder(w.Body).Decode(&stats); err != nil {
		t.Fatal("Problem with decode json:", err)
	}

	if stats.OpenPRAmount != 2 {
		t.Fatalf("waiting 1, get:  %d", stats.OpenPRAmount)
	}
	if stats.MergedPRAmount != 1 {
		t.Fatalf("waiting 1, get:  %d", stats.MergedPRAmount)
	}

}
