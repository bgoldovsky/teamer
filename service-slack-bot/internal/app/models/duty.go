package models

import (
	"fmt"
	"time"
)

type Duty struct {
	TeamID    int64
	PersonID  int64
	FirstName string
	LastName  string
	Slack     string
	Channel   string
	Month     time.Month
	Day       int64
}

func (d *Duty) String() string {
	return fmt.Sprintf("%d %v - %s %s, <@%s>\n",
		d.Day,
		d.Month,
		d.FirstName,
		d.LastName,
		d.Slack)
}
