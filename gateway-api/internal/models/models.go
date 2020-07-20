package models

type Team struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slack       string `json:"slack"`
}
