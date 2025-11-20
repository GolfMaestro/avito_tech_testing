package main

import (
	"avito_tech_testing/repository"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Avito tech testing")

	repository.InitDBConnetion()

	mainController()

	fmt.Println("Servers starts: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
