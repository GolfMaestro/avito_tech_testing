package tests

import (
	"avito_tech_testing/repository"
	"avito_tech_testing/service"
	"context"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
)

func TestCreatePullRequest(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE pull_requests CASCADE;")
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
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE pull_requests CASCADE;")

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

func TestMergePullRequest(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE pull_requests CASCADE;")
	repository.Pool.Exec(context.Background(), "INSERT INTO teams(team_name) VALUES ('t1');")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u124', 'bot2', 't1', false);")
	repository.Pool.Exec(context.Background(), "INSERT INTO pull_requests(pull_request_id, pull_request_name, author_id, assigned_reviewers) VALUES ('pr1', 'update_1', 'u124', '{}');")

	values := strings.NewReader("{\"pull_request_id\": \"pr1\"}")
	req := httptest.NewRequest(http.MethodPost, "/pullRequests/merge", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.MergeRequest(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("waiting 201, get:  %d", w.Code)
	}

	var status string

	repository.Pool.QueryRow(context.Background(), "SELECT status FROM pull_requests WHERE pull_request_id = 'pr1'").Scan(&status)

	if status != "MERGED" {
		t.Fatalf("waiting MERGED, get %s", status)
	}

}

func TestMergePullRequestAlreadyMerged(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE pull_requests CASCADE;")
	repository.Pool.Exec(context.Background(), "INSERT INTO teams(team_name) VALUES ('t1');")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u124', 'bot2', 't1', false);")
	repository.Pool.Exec(context.Background(), "INSERT INTO pull_requests(pull_request_id, pull_request_name, author_id, assigned_reviewers) VALUES ('pr1', 'update_1', 'u124', ARRAY['u1']);")
	repository.Pool.Exec(context.Background(), "UPDATE pull_requests SET status = 'MERGED', merged_at = NOW() WHERE pull_request_id = 'pr1';")

	values := strings.NewReader("{\"pull_request_id\": \"pr1\"}")
	req := httptest.NewRequest(http.MethodPost, "/pullRequests/merge", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.MergeRequest(w, req)

	if w.Code != http.StatusConflict {
		t.Fatalf("waiting 409, get:  %d", w.Code)
	}
}

func TestMergePullRequestPRNotExist(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE pull_requests CASCADE;")
	repository.Pool.Exec(context.Background(), "INSERT INTO teams(team_name) VALUES ('t1');")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u124', 'bot2', 't1', false);")

	values := strings.NewReader("{\"pull_request_id\": \"pr1\"}")
	req := httptest.NewRequest(http.MethodPost, "/pullRequests/merge", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.MergeRequest(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("waiting 404, get:  %d", w.Code)
	}
}

func TestReassignPullRequest(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE pull_requests CASCADE;")
	repository.Pool.Exec(context.Background(), "INSERT INTO teams(team_name) VALUES ('t1');")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u1', 'bot1', 't1', true);")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u2', 'bot2', 't1', true);")

	repository.Pool.Exec(context.Background(),
		"INSERT INTO pull_requests(pull_request_id, pull_request_name, author_id, assigned_reviewers) VALUES ('pr1', 'update_1', 'u2', ARRAY['u1']);")

	values := strings.NewReader("{\"pull_request_id\": \"pr1\", \"old_reviewer_id\": \"u1\"}")
	req := httptest.NewRequest(http.MethodPost, "/pullRequests/reassign", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.ReassignRequest(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("waiting 200, get:  %d", w.Code)
	}

	var assign_reviewers []string

	repository.Pool.QueryRow(context.Background(), "SELECT status FROM pull_requests WHERE pull_request_id = 'pr1'").Scan(&assign_reviewers)

	if slices.Contains(assign_reviewers, "u1") {
		t.Fatalf("user not reassigned")
	} else if slices.Contains(assign_reviewers, "u2") {
		t.Fatalf("author cant be reviewer")
	}

}

func TestReassignPullRequestUserNotExist(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE pull_requests CASCADE;")
	repository.Pool.Exec(context.Background(), "INSERT INTO teams(team_name) VALUES ('t1');")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u1', 'bot1', 't1', true);")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u2', 'bot2', 't1', true);")

	repository.Pool.Exec(context.Background(),
		"INSERT INTO pull_requests(pull_request_id, pull_request_name, author_id, assigned_reviewers) VALUES ('pr1', 'update_1', 'u2', ARRAY['u1']);")

	values := strings.NewReader("{\"pull_request_id\": \"pr1\", \"old_reviewer_id\": \"u5\"}")
	req := httptest.NewRequest(http.MethodPost, "/pullRequests/reassign", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.ReassignRequest(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("waiting 404, get:  %d", w.Code)
	}

}

func TestReassignPullRequestPRNotExist(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE pull_requests CASCADE;")
	repository.Pool.Exec(context.Background(), "INSERT INTO teams(team_name) VALUES ('t1');")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u1', 'bot1', 't1', true);")

	values := strings.NewReader("{\"pull_request_id\": \"pr1\", \"old_reviewer_id\": \"u1\"}")
	req := httptest.NewRequest(http.MethodPost, "/pullRequests/reassign", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.ReassignRequest(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("waiting 404, get:  %d", w.Code)
	}

}

func TestReassignPullRequestPRAlreadyMerged(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE teams CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE;")
	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE pull_requests CASCADE;")
	repository.Pool.Exec(context.Background(), "INSERT INTO teams(team_name) VALUES ('t1');")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u1', 'bot1', 't1', true);")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u2', 'bot2', 't1', true);")

	repository.Pool.Exec(context.Background(),
		"INSERT INTO pull_requests(pull_request_id, pull_request_name, author_id, assigned_reviewers) VALUES ('pr1', 'update_1', 'u2', ARRAY['u1']);")
	repository.Pool.Exec(context.Background(), "UPDATE pull_requests SET status = 'MERGED', merged_at = NOW() WHERE pull_request_id = 'pr1';")

	values := strings.NewReader("{\"pull_request_id\": \"pr1\", \"old_reviewer_id\": \"u1\"}")
	req := httptest.NewRequest(http.MethodPost, "/pullRequests/reassign", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.ReassignRequest(w, req)

	if w.Code != http.StatusConflict {
		t.Fatalf("waiting 409, get:  %d", w.Code)
	}

}
