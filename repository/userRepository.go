package repository

import (
	"context"
	"fmt"
)

func UpdatePersonStatus(req_id string, new_status bool) {

	_, err := Pool.Exec(context.Background(),
		"UPDATE users SET is_active = $1 WHERE user_id = $2", new_status, req_id)
	if err != nil {
		fmt.Println("Something went wrong in funciton UpdatePersonNameById")
	}

	_, err2 := Pool.Exec(context.Background(),
		"UPDATE team_members SET is_active = $1 WHERE user_id = $2", new_status, req_id)
	if err2 != nil {
		fmt.Println("Something went wrong in funciton UpdatePersonNameById")
	}

}
