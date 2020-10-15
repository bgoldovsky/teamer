package persons

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
	personID int64,
	form *models.PersonForm,
	created time.Time,
	updated time.Time,
) *Client {
	m := &ClientMock{}
	m.ConfigureAddPerson(personID, form)
	m.ConfigureGetPersons(personID, form, created, updated)
	m.ConfigureRemovePerson(personID)
	m.ConfigureUpdatePerson(personID, form)
	return newClient(m)
}

func (m *ClientMock) AddPerson(_ context.Context, in *v1.AddPersonRequest, _ ...grpc.CallOption) (*v1.AddPersonReply, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*v1.AddPersonReply)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureAddPerson(
	personID int64,
	form *models.PersonForm,
) {
	arg := &v1.AddPersonRequest{
		TeamId:    form.TeamId,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Slack:     form.Slack,
		Role:      v1.Role(form.Role),
		IsActive:  form.IsActive,
	}

	fake := &v1.AddPersonReply{Id: personID}

	m.On(`AddPerson`, arg).Return(fake, nil)
}

func (m *ClientMock) UpdatePerson(_ context.Context, in *v1.UpdatePersonRequest, _ ...grpc.CallOption) (*empty.Empty, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*empty.Empty)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureUpdatePerson(
	personID int64,
	form *models.PersonForm,
) {
	arg := &v1.UpdatePersonRequest{
		Id:        personID,
		TeamId:    form.TeamId,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Slack:     form.Slack,
		Role:      v1.Role(form.Role),
		IsActive:  form.IsActive,
	}

	fake := &empty.Empty{}

	m.On(`UpdatePerson`, arg).Return(fake, nil)
}

func (m *ClientMock) RemovePerson(_ context.Context, in *v1.RemovePersonRequest, _ ...grpc.CallOption) (*empty.Empty, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*empty.Empty)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureRemovePerson(personID int64) {
	arg := &v1.RemovePersonRequest{Id: personID}
	fake := &empty.Empty{}

	m.On(`RemovePerson`, arg).Return(fake, nil)
}

func (m *ClientMock) GetPersons(_ context.Context, in *v1.GetPersonsRequest, _ ...grpc.CallOption) (*v1.GetPersonsReply, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*v1.GetPersonsReply)
	return res, args.Error(1)
}

func (m *ClientMock) GetPerson(_ context.Context, in *v1.GetPersonRequest, _ ...grpc.CallOption) (*v1.GetPersonReply, error) {
	args := m.Called(in)
	res, _ := args.Get(0).(*v1.GetPersonReply)
	return res, args.Error(1)
}

func (m *ClientMock) ConfigureGetPersons(
	personID int64,
	form *models.PersonForm,
	created time.Time,
	updated time.Time,
) {
	arg := &v1.GetPersonsRequest{
		Filter: nil,
		Limit:  1000,
		Offset: 0,
		Order:  "id",
		Sort:   "desc",
	}

	person := &v1.Person{
		Id:        personID,
		TeamId:    form.TeamId,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Slack:     form.Slack,
		Role:      v1.Role(form.Role),
		IsActive:  form.IsActive,
		Created:   models.ToTimestamp(created),
		Updated:   models.ToTimestamp(updated),
	}

	fakeGetPersons := &v1.GetPersonsReply{Persons: []*v1.Person{person}}
	fakeGetPerson := &v1.GetPersonReply{Person: person}

	m.On(`GetPerson`, arg).Return(fakeGetPerson, nil)
	m.On(`GetPersons`, arg).Return(fakeGetPersons, nil)
}
