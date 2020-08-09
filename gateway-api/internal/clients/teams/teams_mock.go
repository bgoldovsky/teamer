package teams

import (
	"context"
	"time"

	v1 "github.com/bgoldovsky/dutyer/gateway-api/internal/generated/clients/teams/v1"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ClientMock struct {
	mock.Mock
}

func NewMock(
	id int64,
	name string,
	description string,
	slack string,
	created time.Time,
	updated time.Time,
) *Client {
	m := &ClientMock{}
	m.ConfigureAddTeam(id, name, description, slack)
	m.ConfigureGetTeams(id, name, description, slack, created, updated)
	m.ConfigureRemoveTeam(id)
	m.ConfigureUpdateTeam(id, name, description, slack)
	return newClient(m)
}

func (m *ClientMock) AddTeam(_ context.Context, in *v1.AddTeamRequest, _ ...grpc.CallOption) (*v1.AddTeamReply, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*v1.AddTeamReply)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureAddTeam(
	id int64,
	name string,
	description string,
	slack string,
) {
	arg := &v1.AddTeamRequest{
		Name:        name,
		Description: description,
		Slack:       slack,
	}

	fake := &v1.AddTeamReply{Id: id}

	m.On(`AddTeam`, arg).Return(fake, nil)
}

func (m *ClientMock) UpdateTeam(_ context.Context, in *v1.UpdateTeamRequest, _ ...grpc.CallOption) (*empty.Empty, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*empty.Empty)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureUpdateTeam(
	id int64,
	name string,
	description string,
	slack string,
) {
	arg := &v1.UpdateTeamRequest{
		Id:          id,
		Name:        name,
		Description: description,
		Slack:       slack,
	}

	fake := &empty.Empty{}

	m.On(`UpdateTeam`, arg).Return(fake, nil)
}

func (m *ClientMock) RemoveTeam(_ context.Context, in *v1.RemoveTeamRequest, _ ...grpc.CallOption) (*empty.Empty, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*empty.Empty)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureRemoveTeam(id int64) {
	arg := &v1.RemoveTeamRequest{Id: id}
	fake := &empty.Empty{}

	m.On(`RemoveTeam`, arg).Return(fake, nil)
}

func (m *ClientMock) GetTeams(_ context.Context, in *v1.GetTeamsRequest, _ ...grpc.CallOption) (*v1.GetTeamsReply, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*v1.GetTeamsReply)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureGetTeams(
	id int64,
	name string,
	description string,
	slack string,
	created time.Time,
	updated time.Time,
) {
	arg := &v1.GetTeamsRequest{
		Filter: nil,
		Limit:  1000,
		Offset: 0,
		Order:  "id",
		Sort:   "desc",
	}

	fake := &v1.GetTeamsReply{
		Teams: []*v1.Team{
			{
				Id:          id,
				Name:        name,
				Description: description,
				Slack:       slack,
				Created:     models.ToTimestamp(created),
				Updated:     models.ToTimestamp(updated),
			},
		},
	}

	m.On(`GetTeams`, arg).Return(fake, nil)
}
