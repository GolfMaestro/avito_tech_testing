package dto

type Stats struct {
	OpenPRAmount   int `json:"pr_open_amount"`
	MergedPRAmount int `json:"pr_merged_amount"`
}
