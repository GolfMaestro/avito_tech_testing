package repository

import (
	"avito_tech_testing/models"
	"context"
	"fmt"
)

func AddPullRequestInDB(newPullRequest models.PullRequest) {

	reviewers := getNewReviewers(newPullRequest)

	_, err := Pool.Exec(context.Background(),
		"INSERT INTO pull_requests (pull_request_id, pull_request_name, author_id, assigned_reviewers) VALUES ($1, $2, $3, $4);",
		newPullRequest.PullRequestID, newPullRequest.PullRequestName, newPullRequest.AuthorID, reviewers)

	if err != nil {
		fmt.Println("Something went wrong in funciton InsertNewPersonInDB")
	}

}

func getNewReviewers(newPullRequest models.PullRequest) []string {

	var authorTeam string

	err1 := Pool.QueryRow(context.Background(),
		"SELECT team_name FROM users WHERE user_id = $1", newPullRequest.AuthorID).Scan(&authorTeam)

	if err1 != nil {
		fmt.Println("Something went wrong in function GetUsersFromDB")
	}

	rows, err2 := Pool.Query(context.Background(),
		"SELECT user_id FROM users WHERE team_name = $1 AND is_active = true AND user_id != $2;", authorTeam, newPullRequest.AuthorID)

	var userIDs []string

	for rows.Next() {
		var t string
		temp_err := rows.Scan(&t)
		if temp_err != nil {
			fmt.Println("Something went wrong in loop")
		}
		userIDs = append(userIDs, t)
	}

	var reviewers []string

	switch len(userIDs) {
	case 0:
		reviewers = []string{}
	case 1:
		reviewers = []string{userIDs[0]}
	case 2:
		reviewers = []string{userIDs[0], userIDs[1]}
	}

	if err2 != nil {
		fmt.Println("Something went wrong in function GetUsersFromDB")
	}

	return reviewers
}
