package models

import "time"

const (
	RoleNone     = 0
	RoleBackEnd  = 1
	RoleFrontEnd = 2
	RoleMobile   = 3
	RoleQA       = 4
)

type TeamForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slack       string `json:"slack"`
}

type PersonForm struct {
	TeamId     int64      `json:"teamId"`
	FirstName  string     `json:"firstName"`
	MiddleName *string    `json:"middleName"`
	LastName   string     `json:"lastName"`
	Birthday   *time.Time `json:"birthday"`
	Email      *string    `json:"email"`
	Phone      *string    `json:"phone"`
	Slack      string     `json:"slack"`
	Role       int64      `json:"role"`
	IsActive   bool       `json:"isActive"`
}
