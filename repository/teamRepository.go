package repository

import (
	"avito_tech_testing/dto"
	"avito_tech_testing/models"
	"context"
	"errors"
	"fmt"
)

func GetTeamMembersFromDB(teamID string) (models.Team, error) {

	var isExist bool
	err1 := Pool.QueryRow(context.Background(),
		"SELECT EXISTS (SELECT 1 FROM teams WHERE team_name = $1) AS exists;", teamID).Scan(&isExist)

	if err1 != nil {
		fmt.Println("Something went wrong when checking team existence in function GetTeamMembersFromDB")
	}

	if isExist {
		rows, err := Pool.Query(context.Background(),
			"SELECT user_id, username, is_active FROM users WHERE team_name = $1;", teamID)
		if err != nil {
			fmt.Println("Something went wrong when selecting user info in function GetTeamMembersFromDB")
		} else {
			defer rows.Close()

			var team_members []dto.TeamMember

			for rows.Next() {
				var t dto.TeamMember
				temp_err := rows.Scan(&t.UserID, &t.Username, &t.IsActive)
				if temp_err != nil {
					fmt.Println("Something went wrong in parsing team members")
				}
				team_members = append(team_members, t)
			}

			tempTeam := models.Team{
				TeamName: teamID,
				Members:  team_members,
			}

			return tempTeam, nil
		}
	}

	emptyTeam := models.Team{
		TeamName: "0",
		Members:  []dto.TeamMember{},
	}

	return emptyTeam, errors.New("NOT_FOUND")
}

func AddNewTeamToDB(newTeam models.Team) (models.Team, error) {

	var isExist bool
	err2 := Pool.QueryRow(context.Background(),
		"SELECT EXISTS (SELECT 1 FROM teams WHERE team_name = $1) AS exists;", newTeam.TeamName).Scan(&isExist)

	if err2 != nil {
		fmt.Println("Something went wrong when checking team existence in function AddNewTeamToDB")
	}

	if !isExist {
		_, err1 := Pool.Exec(context.Background(),
			"INSERT INTO teams(team_name) VALUES ($1);",
			newTeam.TeamName)

		if err1 != nil {
			fmt.Println("Something went wrong when inserting team in DB in funciton AddNewTeamToDB")
		}

		for _, v := range newTeam.Members {
			var tempMember dto.TeamMember
			tempMember = v

			_, err := Pool.Exec(context.Background(),
				"INSERT INTO users(user_id, username, team_name, is_active) VALUES ($1, $2, $3, $4);",
				tempMember.UserID, tempMember.Username, newTeam.TeamName, tempMember.IsActive)

			if err != nil {
				fmt.Println("Something went wrong when inserting user in DB in funciton AddNewTeamToDB")
			}

		}

		tempTeam := models.Team{
			TeamName: newTeam.TeamName,
			Members:  newTeam.Members,
		}

		return tempTeam, nil

	}

	emptyTeam := models.Team{
		TeamName: "0",
		Members:  []dto.TeamMember{},
	}

	return emptyTeam, errors.New("TEAM_EXISTS")

}
