package main

import (
	"fmt"
	"net/http"
)

func mainController() {

	http.HandleFunc("/main", mainPageHandler)

}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is main page for avito tech testing")
}
