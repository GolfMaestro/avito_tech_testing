package repository

import (
	"avito_tech_testing/models"
	"context"
	"fmt"
	"slices"
)

func AddPullRequestInDB(newPullRequest models.PullRequest) {

	reviewers := getNewReviewers(newPullRequest.AuthorID)

	_, err := Pool.Exec(context.Background(),
		"INSERT INTO pull_requests (pull_request_id, pull_request_name, author_id, assigned_reviewers) VALUES ($1, $2, $3, $4);",
		newPullRequest.PullRequestID, newPullRequest.PullRequestName, newPullRequest.AuthorID, reviewers)

	if err != nil {
		fmt.Println("Something went wrong in funciton InsertNewPersonInDB")
	}

}

func getNewReviewers(author_id string) []string {

	var authorTeam string

	err1 := Pool.QueryRow(context.Background(),
		"SELECT team_name FROM users WHERE user_id = $1", author_id).Scan(&authorTeam)

	if err1 != nil {
		fmt.Println("Something went wrong in function GetUsersFromDB")
	}

	userIDs := getActiveTeamMembersIds(authorTeam, author_id)

	var reviewers []string

	switch len(userIDs) {
	case 0:
		reviewers = []string{}
	case 1:
		reviewers = []string{userIDs[0]}
	case 2:
		reviewers = []string{userIDs[0], userIDs[1]}
	}

	return reviewers
}

func MergeRequestInDB(merger_request_id string) {

	var tempStatus string

	err1 := Pool.QueryRow(context.Background(),
		"SELECT status FROM pull_requests WHERE pull_request_id = $1", merger_request_id).Scan(&tempStatus)

	if err1 != nil {
		fmt.Println("Something went wrong in function MergeRequestInDB")
	}

	if tempStatus == "OPEN" {
		_, err := Pool.Exec(context.Background(),
			"UPDATE pull_requests SET status = 'MERGED', merged_at = NOW() WHERE pull_request_id = $1;", merger_request_id)

		if err != nil {
			fmt.Println("Something went wrong in function MergeRequestInDB")
		}
	}
}

func ReassignRequestInDB(pullRequestID string, oldReviewerID string) {

	var authorTeam string

	var author_id string

	err3 := Pool.QueryRow(context.Background(),
		"SELECT author_id FROM pull_requests WHERE pull_request_id = $1", pullRequestID).Scan(&author_id)
	if err3 != nil {
		fmt.Println("Something went wrong in function err3")
	}

	err1 := Pool.QueryRow(context.Background(),
		"SELECT team_name FROM users WHERE user_id = $1", author_id).Scan(&authorTeam)

	if err1 != nil {
		fmt.Println("Something went wrong in function err1")
	}

	userIDs := getActiveTeamMembersIds(authorTeam, author_id)

	var a_reviewers []string

	err4 := Pool.QueryRow(context.Background(),
		"SELECT assigned_reviewers FROM pull_requests WHERE author_id = $1", author_id).Scan(&a_reviewers)
	if err4 != nil {
		fmt.Println("Something went wrong in function err4")
	}

	var req_id string

	for _, v := range userIDs {
		if !slices.Contains(a_reviewers, v) {
			req_id = v
			break
		}
	}

	for i, v := range a_reviewers {
		if v == oldReviewerID {
			a_reviewers[i] = req_id
			break
		}
	}

	_, err5 := Pool.Exec(context.Background(),
		"UPDATE pull_requests SET assigned_reviewers = $1 WHERE author_id = $2;", a_reviewers, author_id)
	if err5 != nil {
		fmt.Println("Something went wrong in function err5")
	}

}

func getActiveTeamMembersIds(author_team string, author_id string) []string {

	rows, err := Pool.Query(context.Background(),
		"SELECT user_id FROM users WHERE team_name = $1 AND is_active = true AND user_id != $2;", author_team, author_id)

	if err != nil {
		fmt.Println("Something went wrong in function getActiveTeamMembersIds")
	}

	var userIDs []string

	for rows.Next() {
		var t string
		temp_err := rows.Scan(&t)
		if temp_err != nil {
			fmt.Println("Something went wrong in loop in getActiveTeamMembersIds")
		}
		userIDs = append(userIDs, t)
	}

	return userIDs
}
