package repository

import (
	"avito_tech_testing/dto"
	"avito_tech_testing/models"
	"context"
	"errors"
	"fmt"
)

func GetTeamMembersFromDB(teamID string) ([]dto.TeamMember, error) {

	rows, err := Pool.Query(context.Background(),
		"SELECT user_id, username, is_active FROM users WHERE team_name = $1;", teamID)

	if err != nil {
		fmt.Println("Something went wrong in function GetUsersFromDB")
		return nil, errors.New("NOT_FOUND")
	} else {
		defer rows.Close()

		var team_members []dto.TeamMember

		for rows.Next() {
			var t dto.TeamMember
			temp_err := rows.Scan(&t.UserID, &t.Username, &t.IsActive)
			if temp_err != nil {
				fmt.Println("Something went wrong")
			}
			team_members = append(team_members, t)
		}

		return team_members, nil
	}

}

func AddNewTeamToDB(newTeam models.Team) {

	_, err1 := Pool.Exec(context.Background(),
		"INSERT INTO teams(team_name) VALUES ($1);",
		newTeam.TeamName)

	if err1 != nil {
		fmt.Println("Something went wrong in funciton AddNewTeamToDB")
	}

	for _, v := range newTeam.Members {
		var tempMember dto.TeamMember
		tempMember = v

		_, err := Pool.Exec(context.Background(),
			"INSERT INTO users(user_id, username, team_name, is_active) VALUES ($1, $2, $3, $4);",
			tempMember.UserID, tempMember.Username, newTeam.TeamName, tempMember.IsActive)

		if err != nil {
			fmt.Println("Something went wrong in funciton AddNewTeamToDB")
		}

	}

}
