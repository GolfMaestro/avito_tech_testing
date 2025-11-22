package repository

import (
	"avito_tech_testing/dto"
	"avito_tech_testing/models"
	"context"
	"errors"
	"fmt"
)

func UpdateUserStatus(req_id string, new_status bool) (models.User, error) {

	var isExist bool
	err2 := Pool.QueryRow(context.Background(),
		"SELECT EXISTS (SELECT 1 FROM users WHERE user_id = $1) AS exists;", req_id).Scan(&isExist)

	if err2 != nil {
		fmt.Println("Something went wrong in function GetUsersFromDB")
	}

	if isExist {
		var user_id string
		var username string
		var team_name string
		var is_active bool

		err := Pool.QueryRow(context.Background(),
			"UPDATE users SET is_active = $1 WHERE user_id = $2 RETURNING user_id, username, team_name, is_active;", new_status, req_id,
		).Scan(&user_id, &username, &team_name, &is_active)
		if err != nil {
			fmt.Println("Something went wrong in funciton UpdatePersonNameById")
		}

		user := models.User{
			UserID:   user_id,
			Username: username,
			TeamName: team_name,
			IsActive: is_active,
		}

		return user, nil
	}

	emptyUser := models.User{
		UserID:   "",
		Username: "",
		TeamName: "",
		IsActive: false,
	}

	return emptyUser, errors.New("NOT_FOUND")

}

func GetUserReviewsFromDB(user_id string) (dto.UserPullRequests, error) {

	var isExist bool
	err2 := Pool.QueryRow(context.Background(),
		"SELECT EXISTS (SELECT 1 FROM users WHERE user_id = $1) AS exists;", user_id).Scan(&isExist)

	if err2 != nil {
		fmt.Println("Something went wrong in function GetUsersFromDB")
	}

	if isExist {
		rows, err := Pool.Query(context.Background(),
			"SELECT pull_request_id, pull_request_name, author_id, status FROM pull_requests WHERE $1 = ANY(assigned_reviewers);", user_id)

		if err != nil {
			fmt.Println("Something went wrong in function GetUsersFromDB")
		}

		defer rows.Close()

		var pullRequests []dto.PullRequestShort

		for rows.Next() {
			var t dto.PullRequestShort
			temp_err := rows.Scan(&t.PullRequestID, &t.PullRequestName, &t.AuthorID, &t.Status)
			if temp_err != nil {
				fmt.Println("Something went wrong")
			}
			pullRequests = append(pullRequests, t)
		}

		userPullRequests := dto.UserPullRequests{
			UserID:       user_id,
			PullRequests: pullRequests,
		}

		return userPullRequests, nil
	}

	emptyUserPullRequests := dto.UserPullRequests{
		UserID:       "",
		PullRequests: []dto.PullRequestShort{},
	}

	return emptyUserPullRequests, errors.New("NOT_FOUND")

}
