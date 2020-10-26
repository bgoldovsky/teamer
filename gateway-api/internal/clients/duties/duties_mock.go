package duties

import (
	"context"

	v1 "github.com/bgoldovsky/dutyer/gateway-api/internal/generated/clients/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ClientMock struct {
	mock.Mock
}

func NewMock(
	teamID int64,
	firstPersonID int64,
	secondPersonID int64,
	count int64,
	firstName string,
	lastName string,
	slack string,
	channel string,
	order int64,
) *Client {
	m := &ClientMock{}
	m.ConfigureSwap(teamID, firstPersonID, secondPersonID)
	m.ConfigureAssign(teamID, firstPersonID)
	m.ConfigureGetCurrentDuty(teamID, firstPersonID, firstName, lastName, slack, channel, order)
	m.ConfigureGetDuties(teamID, count, firstPersonID, firstName, lastName, slack, channel, order)

	return newClient(m)
}

func (m *ClientMock) GetCurrentDuty(_ context.Context, in *v1.GetCurrentDutyRequest, _ ...grpc.CallOption) (*v1.GetCurrentDutyReply, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*v1.GetCurrentDutyReply)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureGetCurrentDuty(
	teamID int64,
	personID int64,
	firstName string,
	lastName string,
	slack string,
	channel string,
	order int64) {
	arg := &v1.GetCurrentDutyRequest{TeamId: teamID}

	fake := &v1.GetCurrentDutyReply{
		Duty: &v1.Duty{
			TeamId:    teamID,
			PersonId:  personID,
			FirstName: firstName,
			LastName:  lastName,
			Slack:     slack,
			Channel:   channel,
			DutyOrder: order,
		},
	}

	m.On(`GetCurrentDuty`, arg).Return(fake, nil)
}

func (m *ClientMock) GetDuties(_ context.Context, in *v1.GetDutiesRequest, _ ...grpc.CallOption) (*v1.GetDutiesReply, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*v1.GetDutiesReply)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureGetDuties(
	teamID int64,
	count int64,
	personID int64,
	firstName string,
	lastName string,
	slack string,
	channel string,
	order int64,
) {
	arg := &v1.GetDutiesRequest{
		TeamId: teamID,
		Count:  count,
	}

	fake := &v1.GetDutiesReply{
		Duties: []*v1.Duty{
			{
				TeamId:    teamID,
				PersonId:  personID,
				FirstName: firstName,
				LastName:  lastName,
				Slack:     slack,
				Channel:   channel,
				DutyOrder: order,
			},
		},
	}

	m.On(`GetDuties`, arg).Return(fake, nil)
}

func (m *ClientMock) Assign(_ context.Context, in *v1.AssignRequest, _ ...grpc.CallOption) (*empty.Empty, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*empty.Empty)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureAssign(teamID int64, personID int64) {
	arg := &v1.AssignRequest{
		TeamId:   teamID,
		PersonId: personID,
	}

	fake := &v1.AssignRequest{TeamId: teamID, PersonId: personID}

	m.On(`Assign`, arg).Return(fake, nil)
}

func (m *ClientMock) Swap(_ context.Context, in *v1.SwapRequest, _ ...grpc.CallOption) (*empty.Empty, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*empty.Empty)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureSwap(teamID int64, firstPersonID int64, secondPersonID int64) {
	arg := &v1.SwapRequest{
		TeamId:         teamID,
		FirstPersonId:  firstPersonID,
		SecondPersonId: secondPersonID,
	}

	fake := &v1.AddPersonReply{Id: teamID}

	m.On(`Swap`, arg).Return(fake, nil)
}
