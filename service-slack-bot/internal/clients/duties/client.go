//go:generate mockgen -destination client_mock.go -source client.go

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

type Client interface {
	GetCurrentDuty(ctx context.Context, teamID int64) (*models.Duty, error)
}

type client struct {
	grpcClient v1.DutiesClient
}

func NewClient(host string) (*client, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptors.LoggingInterceptor))
	if err != nil {
		logger.Log.WithError(err).Fatal("can't connect service-dutyer")
	}

	grpcClient := v1.NewDutiesClient(conn)
	if err != nil {
		return nil, err
	}
	return &client{grpcClient: grpcClient}, nil
}

func (c *client) GetCurrentDuty(ctx context.Context, teamID int64) (*models.Duty, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.GetCurrentDutyRequest{TeamId: teamID}

	reply, err := c.grpcClient.GetCurrentDuty(ctx, request)
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
