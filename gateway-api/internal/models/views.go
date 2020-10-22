package models

import (
	"time"

	v1 "github.com/bgoldovsky/dutyer/gateway-api/internal/generated/clients/v1"
)

type StatusView struct {
	ID  int64  `json:"id"`
	Msg string `json:"message"`
}

func NewStatusView(id int64, msg string) *StatusView {
	return &StatusView{
		ID:  id,
		Msg: msg,
	}
}

type TeamView struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Slack       string    `json:"slack"`
	Crated      time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

func FromTeamsReply(reply *v1.GetTeamsReply) []TeamView {
	if reply == nil || len(reply.Teams) == 0 {
		return []TeamView{}
	}

	view := make([]TeamView, len(reply.Teams))
	for idx, t := range reply.Teams {
		team := FromTeamReply(t)
		view[idx] = *team
	}
	return view
}

func FromTeamReply(reply *v1.Team) *TeamView {
	if reply == nil {
		return nil
	}

	return &TeamView{
		ID:          reply.Id,
		Name:        reply.Name,
		Description: reply.Description,
		Slack:       reply.Slack,
		Crated:      ToTime(reply.Created),
		Updated:     ToTime(reply.Updated),
	}
}

type PersonView struct {
	ID         int64      `json:"id"`
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
	Created    time.Time  `json:"created"`
	Updated    time.Time  `json:"updated"`
}

func FromPersonsReply(reply *v1.GetPersonsReply) []PersonView {
	if reply == nil || len(reply.Persons) == 0 {
		return []PersonView{}
	}

	view := make([]PersonView, len(reply.Persons))
	for idx, p := range reply.Persons {
		person := FromPersonReply(p)
		view[idx] = *person
	}
	return view
}

func FromPersonReply(reply *v1.Person) *PersonView {
	if reply == nil {
		return nil
	}

	view := PersonView{
		ID:        reply.Id,
		TeamId:    reply.TeamId,
		FirstName: reply.FirstName,
		LastName:  reply.LastName,
		Slack:     reply.Slack,
		Role:      int64(reply.Role),
		IsActive:  reply.IsActive,
		Created:   ToTime(reply.Created),
		Updated:   ToTime(reply.Updated),
	}

	if reply.MiddleName != nil {
		view.MiddleName = &reply.MiddleName.Value
	}

	if reply.Birthday != nil {
		birthday := ToTime(reply.Birthday)
		view.Birthday = &birthday
	}

	if reply.Email != nil {
		view.Email = &reply.Email.Value
	}

	if reply.Phone != nil {
		view.Phone = &reply.Phone.Value
	}

	return &view
}
