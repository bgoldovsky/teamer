package models

import "time"

type Duty struct {
	TeamID    int64      `json:"teamId"`
	PersonID  int64      `json:"personId"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Slack     string     `json:"slack"`
	Order     int64      `json:"order"`
	Month     time.Month `json:"month"`
	Day       int64      `json:"day"`
	Created   time.Time  `json:"created"`
	Updated   time.Time  `json:"updated"`
}
