package service

import (
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
