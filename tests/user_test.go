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

func TestUpdateUserStatus(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE")
	repository.Pool.Exec(context.Background(), "INSERT INTO users(user_id, username, team_name, is_active) VALUES ('u1', 'bot1', 't1', false);")

	values := strings.NewReader("{\"user_id\": \"u1\", \"is_active\": true}")
	req := httptest.NewRequest(http.MethodPost, "/users/setIsActive", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.UpdateUserStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("waiting 200, get:  %d", w.Code)
	}

	var is_active bool
	repository.Pool.QueryRow(context.Background(), "SELECT is_active FROM users WHERE user_id = 'u1'").Scan(&is_active)
	if !is_active {
		t.Fatal("User is_active = false, but waiting true")
	}

}

func TestUpdateUserStatusUserNotFound(t *testing.T) {

	TestConnection(t)

	repository.Pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE")

	values := strings.NewReader("{\"user_id\": \"u1\", \"is_active\": true}")
	req := httptest.NewRequest(http.MethodPost, "/users/setIsActive", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.UpdateUserStatus(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("waiting 404, get:  %d", w.Code)
	}

}
