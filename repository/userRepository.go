package repository

import (
	"avito_tech_testing/dto"
	"context"
	"fmt"
)

func UpdateUserStatus(req_id string, new_status bool) {

	_, err := Pool.Exec(context.Background(),
		"UPDATE users SET is_active = $1 WHERE user_id = $2", new_status, req_id)
	if err != nil {
		fmt.Println("Something went wrong in funciton UpdatePersonNameById")
	}

}

func GetUserReviewsFromDB(user_id string) []dto.PullRequestShort {

	rows, err := Pool.Query(context.Background(),
		"SELECT pull_request_id, pull_request_name, author_id, status FROM pull_requests WHERE $1 = ANY(assigned_reviewers);", user_id)

	if err != nil {
		fmt.Println("Something went wrong in function GetUsersFromDB")
	}

	defer rows.Close()

	var team_members []dto.PullRequestShort

	for rows.Next() {
		var t dto.PullRequestShort
		temp_err := rows.Scan(&t.PullRequestID, &t.PullRequestName, &t.AuthorID, &t.Status)
		if temp_err != nil {
			fmt.Println("Something went wrong")
		}
		team_members = append(team_members, t)
	}

	return team_members

}
