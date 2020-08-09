package models

import (
	"time"

	v1 "github.com/bgoldovsky/dutyer/gateway-api/internal/generated/clients/teams/v1"
)

type TeamView struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Slack       string    `json:"slack"`
	Crated      time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

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

func FromTeamsReply(reply *v1.GetTeamsReply) []*TeamView {
	if reply == nil || len(reply.Teams) == 0 {
		return []*TeamView{}
	}

	view := make([]*TeamView, len(reply.Teams))
	for idx, t := range reply.Teams {
		view[idx] = FromTeamReply(t)
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
