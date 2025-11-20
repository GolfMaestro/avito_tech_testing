package repository

import (
	"avito_tech_testing/models"
	"context"
	"fmt"
)

func AddPullRequestInDB(newPullRequest models.PullRequest) {

	_, err := Pool.Exec(context.Background(),
		"INSERT INTO pull_requests(pull_request_id, pull_request_name, author_id, status) VALUES ($1, $2, $3, 'OPEN');",
		newPullRequest.PullRequestID, newPullRequest.PullRequestName, newPullRequest.AuthorID)

	if err != nil {
		fmt.Println("Something went wrong in funciton InsertNewPersonInDB")
	}

}
