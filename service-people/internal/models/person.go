package models

import "time"

type Role int64

const (
	None Role = iota
	BackEnd
	FrontEnd
	Mobile
	QA
)

type Person struct {
	ID         int64
	FirstName  string
	MiddleName *string
	LastName   string
	Birthday   int64
	Email      *string
	Phone      *string
	Slack      string
	Role       Role
	TeamID     *int64
	IsActive   bool
	Created    time.Time
	Updated    time.Time
}
