package repository

import (
	"avito_tech_testing/dto"
	"context"
	"fmt"
)

func GetAllPRStats() dto.Stats {

	var openPR int
	var mergedPR int

	err1 := Pool.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM pull_requests WHERE status = 'OPEN';").Scan(&openPR)
	if err1 != nil {
		fmt.Println("Something went wrong when collecting open PR in function GetAllPRStats")
	}

	err2 := Pool.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM pull_requests WHERE status = 'MERGED';").Scan(&mergedPR)
	if err2 != nil {
		fmt.Println("Something went wrong when collecting merged PR in function GetAllPRStats")
	}

	stats := dto.Stats{
		OpenPRAmount:   openPR,
		MergedPRAmount: mergedPR,
	}

	return stats

}
