package main

import (
	"avito_tech_testing/service"
	"fmt"
	"net/http"
)

func mainController() {

	http.HandleFunc("/main", mainPageHandler)

	http.HandleFunc("/users/setIsActive", service.UpdateUserStatus)
	http.HandleFunc("/users/getReview", service.GetUserReviews)

	http.HandleFunc("/pullRequest/create", service.CreatePullRequest)
	http.HandleFunc("/pullRequest/merge", service.MergeRequest)
	http.HandleFunc("/pullRequest/reassign", service.ReassignRequest)

	http.HandleFunc("/team/get", service.GetTeamMembers)
	http.HandleFunc("/team/add", service.CreateNewTeam)

}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is main page for avito tech testing")
}
