package dto

type ReassignPullRequestResponse struct {
	PR         CreatePullRequestResponse `json:"pr"`
	ReplacedBy string                    `json:"replaced_by"`
}
