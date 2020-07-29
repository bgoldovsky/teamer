package teams

import (
	"context"

	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/clients/teams"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/logger"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/models"
	teamsRepo "github.com/bgoldovsky/teamer-bot/gateway-api/internal/repostiory/teams"
	"github.com/sirupsen/logrus"
)

type Service struct {
	client *teams.Client
	repo   teamsRepo.Repository
}

func New(client *teams.Client, repo teamsRepo.Repository) *Service {
	return &Service{
		client: client,
		repo:   repo,
	}
}

func (s *Service) AddTeam(ctx context.Context, form *models.TeamForm) (*models.StatusView, error) {
	status, err := s.client.AddTeam(ctx, form.Name, form.Description, form.Slack)
	if err != nil {
		logger.Log.WithField("form", form).WithError(err).Errorln("add team error")
		return nil, err
	}

	s.clearRepo()
	return status, nil
}

func (s *Service) UpdateTeam(ctx context.Context, id int64, form *models.TeamForm) (*models.StatusView, error) {
	status, err := s.client.UpdateTeam(ctx, id, form.Name, form.Description, form.Slack)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"form": form, "id": id}).WithError(err).Errorln("update team error")
		return nil, err
	}

	s.clearRepo()
	return status, nil
}

func (s *Service) RemoveTeam(ctx context.Context, id int64) (*models.StatusView, error) {
	status, err := s.client.RemoveTeam(ctx, id)
	if err != nil {
		logger.Log.WithField("id", id).WithError(err).Errorln("remove team error")
		return nil, err
	}

	s.clearRepo()
	return status, nil
}

func (s *Service) GetTeams(ctx context.Context) ([]*models.TeamView, error) {
	if view := s.getRepo(); view != nil {
		return view, nil
	}

	view, err := s.client.GetTeams(ctx)
	if err != nil {
		logger.Log.WithError(err).Errorln("get teams error")
		return nil, err
	}

	s.saveRepo(view)
	return view, nil
}

func (s *Service) clearRepo() {
	err := s.repo.Clear()
	if err != nil {
		logger.Log.WithError(err).Error("clear teams repo error")
	}
}

func (s *Service) saveRepo(teams []*models.TeamView) {
	err := s.repo.Save(teams)
	if err != nil {
		logger.Log.WithError(err).Error("save teams repo error")
	}
}

func (s *Service) getRepo() []*models.TeamView {
	views, err := s.repo.Get()
	if err != nil {
		logger.Log.WithError(err).Error("save teams repo error")
		return nil
	}
	return views
}
