package teams

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bgoldovsky/dutyer/service-teams/internal/app/services/teams"
	v1 "github.com/bgoldovsky/dutyer/service-teams/internal/generated/rpc/v1"
	"github.com/bgoldovsky/dutyer/service-teams/internal/logger"
	"github.com/golang/protobuf/ptypes/empty"
)

type Handler struct {
	service *teams.Service
}

func New(service *teams.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddTeam(ctx context.Context, req *v1.AddTeamRequest) (*v1.AddTeamReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "team name not specified")
	}

	if len(req.Slack) == 0 {
		return nil, status.Error(codes.InvalidArgument, "team slack not specified")
	}

	reply, err := h.service.AddTeam(ctx, req)
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Errorln("add team error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return reply, nil
}

func (h *Handler) UpdateTeam(ctx context.Context, req *v1.UpdateTeamRequest) (*empty.Empty, error) {
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "team name not specified")
	}

	if len(req.Slack) == 0 {
		return nil, status.Error(codes.InvalidArgument, "team slack not specified")
	}

	reply, err := h.service.UpdateTeam(ctx, req)
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Errorln("update team error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return reply, nil
}

func (h *Handler) RemoveTeam(ctx context.Context, req *v1.RemoveTeamRequest) (*empty.Empty, error) {
	reply, err := h.service.RemoveTeam(ctx, req)
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Errorln("remove team error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return reply, nil
}

func (h *Handler) GetTeams(ctx context.Context, req *v1.GetTeamsRequest) (*v1.GetTeamsReply, error) {
	if req.Order != "id" && req.Order != "name" {
		return nil, status.Error(codes.InvalidArgument, "order must be id|name")
	}

	if req.Sort != "asc" && req.Sort != "desc" {
		return nil, status.Error(codes.InvalidArgument, "sort must be asc|desc")
	}

	reply, err := h.service.GetTeams(ctx, req)
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Errorln("get teams error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return reply, nil
}

func (h *Handler) GetTeam(context.Context, *v1.GetTeamRequest) (*v1.GetTeamReply, error) {
	return nil, errors.New("not implemented")
}
