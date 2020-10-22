package teams

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
	GetTeam(ctx context.Context, in *v1.GetTeamRequest, opts ...grpc.CallOption) (*v1.GetTeamReply, error)
	GetTeams(ctx context.Context, in *v1.GetTeamsRequest, opts ...grpc.CallOption) (*v1.GetTeamsReply, error)
	AddTeam(ctx context.Context, in *v1.AddTeamRequest, opts ...grpc.CallOption) (*v1.AddTeamReply, error)
	UpdateTeam(ctx context.Context, in *v1.UpdateTeamRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveTeam(ctx context.Context, in *v1.RemoveTeamRequest, opts ...grpc.CallOption) (*empty.Empty, error)
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

	client := v1.NewTeamsClient(conn)
	if err != nil {
		return nil, err
	}
	return newClient(client), nil
}

func (c *Client) AddTeam(ctx context.Context, name, description, slack string) (*models.StatusView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.AddTeamRequest{
		Name:        name,
		Description: description,
		Slack:       slack,
	}

	reply, err := c.client.AddTeam(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewStatusView(reply.Id, "successfully created"), nil
}

func (c *Client) UpdateTeam(ctx context.Context, teamID int64, name, description, slack string) (*models.StatusView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.UpdateTeamRequest{
		Id:          teamID,
		Name:        name,
		Description: description,
		Slack:       slack,
	}

	_, err := c.client.UpdateTeam(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewStatusView(teamID, "successfully updated"), nil
}

func (c *Client) RemoveTeam(ctx context.Context, teamID int64) (*models.StatusView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.RemoveTeamRequest{Id: teamID}

	_, err := c.client.RemoveTeam(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewStatusView(teamID, "successfully removed"), nil
}

func (c *Client) GetTeam(ctx context.Context, teamID int64) (*models.TeamView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.GetTeamRequest{Id: teamID}

	team, err := c.client.GetTeam(ctx, request)
	if err != nil {
		return nil, err
	}

	view := models.FromTeamReply(team.Team)
	return view, nil
}

func (c *Client) GetTeams(ctx context.Context) ([]models.TeamView, error) {
	ctx = getTimeoutContext(ctx)
	request := &v1.GetTeamsRequest{
		Limit:  1000,
		Offset: 0,
		Order:  "id",
		Sort:   "desc",
	}

	teams, err := c.client.GetTeams(ctx, request)
	if err != nil {
		return nil, err
	}

	view := models.FromTeamsReply(teams)
	return view, nil
}

func getTimeoutContext(ctx context.Context) context.Context {
	ctx, _ = context.WithTimeout(ctx, time.Second*timeout)
	return ctx
}
