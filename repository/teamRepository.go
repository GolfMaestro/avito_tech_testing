package repository

import (
	"avito_tech_testing/models"
	"context"
	"fmt"
)

func GetTeamMembersFromDB(teamID string) []models.TeamMember {

	rows, err := Pool.Query(context.Background(),
		"SELECT user_id, username, is_active FROM users WHERE team_name = $1;", teamID)

	if err != nil {
		fmt.Println("Something went wrong in function GetUsersFromDB")
	}

	defer rows.Close()

	var team_members []models.TeamMember

	for rows.Next() {
		var t models.TeamMember
		temp_err := rows.Scan(&t.UserID, &t.Username, &t.IsActive)
		if temp_err != nil {
			fmt.Println("Something went wrong")
		}
		team_members = append(team_members, t)
	}

	return team_members
}
