package main

import (
	"avito_tech_testing/service"
	"fmt"
	"net/http"
)

func mainController() {

	http.HandleFunc("/main", mainPageHandler)

	http.HandleFunc("/users/setIsActive", service.UpdatePersonStatus)

}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is main page for avito tech testing")
}
