package duties

import (
	"context"
	"time"

	v1 "github.com/bgoldovsky/dutyer/gateway-api/internal/generated/clients/v1"
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
	GetCurrentDuty(ctx context.Context, in *v1.GetCurrentDutyRequest, opts ...grpc.CallOption) (*v1.GetCurrentDutyReply, error)
	GetDuties(ctx context.Context, in *v1.GetDutiesRequest, opts ...grpc.CallOption) (*v1.GetDutiesReply, error)
	Assign(ctx context.Context, in *v1.AssignRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Swap(ctx context.Context, in *v1.SwapRequest, opts ...grpc.CallOption) (*empty.Empty, error)
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

	client := v1.NewDutiesClient(conn)
	if err != nil {
		return nil, err
	}
	return newClient(client), nil
}

func (c *Client) GetCurrentDuty(ctx context.Context, teamID int64) (*models.DutyView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.GetCurrentDutyRequest{TeamId: teamID}

	reply, err := c.client.GetCurrentDuty(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.FromDutyReply(reply.Duty), nil
}

func (c *Client) GetDuties(ctx context.Context, teamID int64, count int64) ([]models.DutyView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.GetDutiesRequest{
		TeamId: teamID,
		Count:  count,
	}

	reply, err := c.client.GetDuties(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.FromDutiesReply(reply), nil
}

func (c *Client) Assign(ctx context.Context, teamID int64, personID int64) (*models.StatusView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.AssignRequest{
		TeamId:   teamID,
		PersonId: personID,
	}

	_, err := c.client.Assign(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewStatusView(teamID, "successfully assigned"), nil
}

func (c *Client) Swap(ctx context.Context, teamID, firstPersonID, secondPersonID int64) (*models.StatusView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.SwapRequest{
		TeamId:         teamID,
		FirstPersonId:  firstPersonID,
		SecondPersonId: secondPersonID,
	}

	_, err := c.client.Swap(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewStatusView(teamID, "successfully swapped"), nil
}

func getTimeoutContext(ctx context.Context) context.Context {
	ctx, _ = context.WithTimeout(ctx, time.Second*timeout)
	return ctx
}
