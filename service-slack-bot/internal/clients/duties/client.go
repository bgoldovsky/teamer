//go:generate mockgen -destination duties_mock/duties_mock.go -source client.go

package duties

import (
	"context"
	"time"

	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/app/models"
	v1 "github.com/bgoldovsky/dutyer/service-slack-bot/internal/generated/clients/v1"
	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/interceptors"
	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/logger"
	"google.golang.org/grpc"
)

const (
	timeout = time.Second * 3
)

type client interface {
	GetCurrentDuty(ctx context.Context, in *v1.GetCurrentDutyRequest, opts ...grpc.CallOption) (*v1.GetCurrentDutyReply, error)
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

func (c *Client) GetCurrentDuty(ctx context.Context, teamID int64) (*models.Duty, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.GetCurrentDutyRequest{TeamId: teamID}

	reply, err := c.client.GetCurrentDuty(ctx, request)
	if err != nil {
		return nil, err
	}

	return &models.Duty{
		TeamID:    reply.Duty.TeamId,
		PersonID:  reply.Duty.PersonId,
		FirstName: reply.Duty.FirstName,
		LastName:  reply.Duty.LastName,
		Slack:     reply.Duty.Slack,
		Channel:   reply.Duty.Channel,
		Month:     time.Month(reply.Duty.Month),
		Day:       reply.Duty.Day,
	}, nil
}

func getTimeoutContext(ctx context.Context) context.Context {
	ctx, _ = context.WithTimeout(ctx, time.Second*timeout)
	return ctx
}
