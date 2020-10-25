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
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Slack       string `json:"slack" validate:"required"`
}

type PersonForm struct {
	TeamId     int64      `json:"teamId" validate:"required"`
	FirstName  string     `json:"firstName" validate:"required"`
	MiddleName *string    `json:"middleName"`
	LastName   string     `json:"lastName" validate:"required"`
	Birthday   *time.Time `json:"birthday"`
	Email      *string    `json:"email"`
	Phone      *string    `json:"phone"`
	Slack      string     `json:"slack" validate:"required"`
	Role       int64      `json:"role" validate:"required,gte=0,lte=5"`
	IsActive   bool       `json:"isActive"`
}

type SwapForm struct {
	TeamId         int64 `json:"teamId" validate:"required"`
	FirstPersonID  int64 `json:"firstPersonId" validate:"required"`
	SecondPersonID int64 `json:"secondPersonId" validate:"required"`
}

type AssignForm struct {
	TeamId   int64 `json:"teamId" validate:"required"`
	PersonID int64 `json:"personId" validate:"required"`
}
