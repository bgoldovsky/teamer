package models

import "time"

type Team struct {
	ID          int64
	Name        string
	Description string
	Slack       string
	Created     time.Time
	Updated     time.Time
}
