package persons

import (
	"context"
	"time"

	v1 "github.com/bgoldovsky/dutyer/gateway-api/internal/generated/clients/v1"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/interceptors"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
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

func (c *Client) AddPerson(ctx context.Context, form *models.PersonForm) (*models.StatusView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.AddPersonRequest{
		TeamId:    form.TeamId,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Slack:     form.Slack,
		Role:      v1.Role(form.Role),
		IsActive:  form.IsActive,
	}

	if form.MiddleName != nil {
		request.MiddleName = &wrappers.StringValue{Value: *form.MiddleName}
	}

	if form.Birthday != nil {
		request.Birthday = models.ToTimestamp(*form.Birthday)
	}

	if form.Email != nil {
		request.Email = &wrappers.StringValue{Value: *form.Email}
	}

	if form.Phone != nil {
		request.Phone = &wrappers.StringValue{Value: *form.Phone}
	}

	reply, err := c.client.AddPerson(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewStatusView(reply.Id, "successfully created"), nil
}

func (c *Client) UpdatePerson(ctx context.Context, personID int64, form *models.PersonForm) (*models.StatusView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.UpdatePersonRequest{
		Id:        personID,
		TeamId:    form.TeamId,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Slack:     form.Slack,
		Role:      v1.Role(form.Role),
		IsActive:  form.IsActive,
	}

	if form.MiddleName != nil {
		request.MiddleName = &wrappers.StringValue{Value: *form.MiddleName}
	}

	if form.Birthday != nil {
		request.Birthday = models.ToTimestamp(*form.Birthday)
	}

	if form.Email != nil {
		request.Email = &wrappers.StringValue{Value: *form.Email}
	}

	if form.Phone != nil {
		request.Phone = &wrappers.StringValue{Value: *form.Phone}
	}

	_, err := c.client.UpdatePerson(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewStatusView(personID, "successfully updated"), nil
}

func (c *Client) RemovePerson(ctx context.Context, personID int64) (*models.StatusView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.RemovePersonRequest{Id: personID}

	_, err := c.client.RemovePerson(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewStatusView(personID, "successfully removed"), nil
}

func (c *Client) GetPerson(ctx context.Context, personID int64) (*models.PersonView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.GetPersonRequest{Id: personID}

	person, err := c.client.GetPerson(ctx, request)
	if err != nil {
		return nil, err
	}

	view := models.FromPersonReply(person.Person)
	return view, nil
}

func (c *Client) GetPersons(ctx context.Context, teamID *int64) ([]models.PersonView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.GetPersonsRequest{
		Limit:  1000,
		Offset: 0,
		Order:  "id",
		Sort:   "desc",
	}

	if teamID != nil {
		request.Filter = &v1.PersonFilter{
			TeamIds: []int64{*teamID},
		}
	}

	persons, err := c.client.GetPersons(ctx, request)
	if err != nil {
		return nil, err
	}

	view := models.FromPersonsReply(persons)
	return view, nil
}

func getTimeoutContext(ctx context.Context) context.Context {
	ctx, _ = context.WithTimeout(ctx, time.Second*timeout)
	return ctx
}
