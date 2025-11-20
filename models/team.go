package models

import "avito_tech_testing/dto"

type Team struct {
	TeamName string `json:"team_name"`
	Members  []dto.TeamMember
}
