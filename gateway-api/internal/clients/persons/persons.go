package persons

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"

	v1 "github.com/bgoldovsky/dutyer/gateway-api/internal/generated/clients/teams/v1"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/interceptors"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	timeout = time.Second * 3
)

type client interface {
	GetPerson(ctx context.Context, in *v1.GetPersonRequest, opts ...grpc.CallOption) (*v1.GetPersonReply, error)
	GetPersons(ctx context.Context, in *v1.GetPersonsRequest, opts ...grpc.CallOption) (*v1.GetPersonsReply, error)
	AddPerson(ctx context.Context, in *v1.AddPersonRequest, opts ...grpc.CallOption) (*v1.AddPersonReply, error)
	UpdatePerson(ctx context.Context, in *v1.UpdatePersonRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemovePerson(ctx context.Context, in *v1.RemovePersonRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type Client struct {
	client client
}

func newClient(client client) *Client {
	return &Client{
		client: client,
	}
}

func NewClient(host string) (*Client, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptors.LoggingInterceptor))
	if err != nil {
		logger.Log.WithError(err).Fatal("can't connect service-dutyer")
	}

	client := v1.NewPersonsClient(conn)
	if err != nil {
		return nil, err
	}
	return newClient(client), nil
}

// TODO: Заменить на PersonForm
func (c *Client) AddPerson(
	ctx context.Context,
	teamID int64,
	firstName string,
	middleName *string,
	lastName string,
	birthday *time.Time,
	email *string,
	phone *string,
	slack string,
	role int64,
	isActive bool,
) (*models.StatusView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.AddPersonRequest{
		TeamId:    teamID,
		FirstName: firstName,
		LastName:  lastName,
		Birthday:  nil,
		Email:     nil,
		Phone:     nil,
		Slack:     slack,
		Role:      v1.Role(role),
		IsActive:  isActive,
	}

	if middleName != nil {
		request.MiddleName = &wrappers.StringValue{Value: *middleName}
	}

	if birthday != nil {
		request.Birthday = models.ToTimestamp(*birthday)
	}

	if email != nil {
		request.Email = &wrappers.StringValue{Value: *email}
	}

	if phone != nil {
		request.Phone = &wrappers.StringValue{Value: *phone}
	}

	reply, err := c.client.AddPerson(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewStatusView(reply.Id, "successfully created"), nil
}

// TODO: Реализовать
/*
func (c *Client) UpdatePerson(ctx context.Context, id int64, name, description, slack string) (*models.StatusView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.UpdatePersonRequest{
		Id:          id,
		Name:        name,
		Description: description,
		Slack:       slack,
	}

	_, err := c.client.UpdatePerson(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewStatusView(id, "successfully updated"), nil
}
*/

func (c *Client) RemovePerson(ctx context.Context, id int64) (*models.StatusView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.RemovePersonRequest{
		Id: id,
	}

	_, err := c.client.RemovePerson(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewStatusView(id, "successfully removed"), nil
}

/*
func (c *Client) GetPerson(ctx context.Context, teamID int64) (*models.TeamView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.GetPersonRequest{Id: teamID}

	team, err := c.client.GetPerson(ctx, request)
	if err != nil {
		return nil, err
	}

	view := models.FromPersonReply(person.Person)
	return view, nil
}

func (c *Client) GetPersons(ctx context.Context) ([]*models.PersonView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.GetPersonsRequest{
		Limit:  1000,
		Offset: 0,
		Order:  "id",
		Sort:   "desc",
	}

	persons, err := c.client.GetPersons(ctx, request)
	if err != nil {
		return nil, err
	}

	view := models.FromPersonsReply(persons)
	return view, nil
}
*/

func getTimeoutContext(ctx context.Context) context.Context {
	ctx, _ = context.WithTimeout(ctx, time.Second*timeout)
	return ctx
}
