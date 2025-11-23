package tests

import (
	"avito_tech_testing/repository"
	"avito_tech_testing/service"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreatePullRequest(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE;")
	repository.Pool.Exec(context.Background(), "INSERT INTO teams(team_name) VALUES ('t1');")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u2', 'bot1', 't1', false);")

	values := strings.NewReader("{\"pull_request_id\": \"pr1\", \"pull_request_name\": \"update_1\", \"author_id\": \"u2\"}")
	req := httptest.NewRequest(http.MethodPost, "/pullRequests/create", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.CreatePullRequest(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("waiting 201, get:  %d", w.Code)
	}

	var assignedReviewers []string
	repository.Pool.QueryRow(context.Background(), "SELECT assigned_reviewers FROM pull_requests WHERE author_id = 'u2'").Scan(&assignedReviewers)

	if len(assignedReviewers) != 0 {
		t.Fatalf("0, get:  %d", len(assignedReviewers))
	}

}

func TestCreatePullRequestAuthorNotExist(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE;")

	values := strings.NewReader("{\"pull_request_id\": \"pr1\", \"pull_request_name\": \"update_1\", \"author_id\": \"u2\"}")
	req := httptest.NewRequest(http.MethodPost, "/pullRequests/create", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.CreatePullRequest(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("waiting 404, get:  %d", w.Code)
	}

}

func TestCreatePullRequestPRAlreadyExist(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE pull_requests CASCADE;")
	repository.Pool.Exec(context.Background(), "INSERT INTO teams(team_name) VALUES ('t1');")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u2', 'bot1', 't1', false);")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u124', 'bot2', 't1', false);")
	repository.Pool.Exec(context.Background(), "INSERT INTO pull_requests(pull_request_id, pull_request_name, author_id, assigned_reviewers) VALUES ('pr1', 'update_1', 'u124', '{}');")

	values := strings.NewReader("{\"pull_request_id\": \"pr1\", \"pull_request_name\": \"update_1\", \"author_id\": \"u2\"}")
	req := httptest.NewRequest(http.MethodPost, "/pullRequests/create", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.CreatePullRequest(w, req)

	if w.Code != http.StatusConflict {
		t.Fatalf("waiting 409, get:  %d", w.Code)
	}

}
