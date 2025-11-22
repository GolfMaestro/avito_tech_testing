package utility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func getRequestedId(r *http.Request) int {

	path := r.URL.Path
	parts := strings.Split(path, "/")

	requested_id, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Something went wrong")
	}

	return requested_id
}

func Err(w http.ResponseWriter, status int, code string, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Code: code, Message: msg})
}
