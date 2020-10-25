package teams

import (
	"context"
	"fmt"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/teams"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/services/teams/converter"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

type Service struct {
	repo teams.Repository
}

func New(repo teams.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetTeam(ctx context.Context, req *v1.GetTeamRequest) (*v1.GetTeamReply, error) {
	team, err := s.repo.Get(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("get team error: %w", err)
	}
	return &v1.GetTeamReply{Team: converter.ToDTO(team)}, nil
}

func (s *Service) GetTeams(ctx context.Context, req *v1.GetTeamsRequest) (*v1.GetTeamsReply, error) {
	teamModels, err := s.repo.GetList(ctx, req.Filter, uint(req.Limit), uint(req.Offset), req.Sort, req.Order)
	if err != nil {
		return nil, fmt.Errorf("get teams error: %w", err)
	}

	if len(teamModels) == 0 {
		return &v1.GetTeamsReply{Teams: []*v1.Team{}}, nil
	}

	teamsProto := make([]*v1.Team, len(teamModels))
	for idx, team := range teamModels {
		teamsProto[idx] = converter.ToDTO(&team)
	}

	return &v1.GetTeamsReply{Teams: teamsProto}, nil
}

func (s *Service) AddTeam(ctx context.Context, req *v1.AddTeamRequest) (*v1.AddTeamReply, error) {
	team, err := s.repo.Save(ctx, &models.Team{
		Name:        req.Name,
		Description: req.Description,
		Slack:       req.Slack,
	})

	if err != nil {
		return nil, fmt.Errorf("add team error: %w", err)
	}

	return &v1.AddTeamReply{Id: team.ID}, nil
}

func (s *Service) UpdateTeam(ctx context.Context, req *v1.UpdateTeamRequest) (*empty.Empty, error) {
	team := converter.ToModel(req)

	_, err := s.repo.Update(ctx, team)
	if err != nil {
		return nil, fmt.Errorf("update team error: %w", err)
	}

	return &empty.Empty{}, nil
}

func (s *Service) RemoveTeam(ctx context.Context, req *v1.RemoveTeamRequest) (*empty.Empty, error) {
	_, err := s.repo.Remove(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("remove team error: %w", err)
	}
	return &empty.Empty{}, nil
}
