package duties

import (
	"context"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/duties"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
)

type Service struct {
	client *duties.Client
}

func New(client *duties.Client) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) Swap(ctx context.Context, teamID, firstPersonID, secondPersonID int64) (*models.StatusView, error) {
	view, err := s.client.Swap(ctx, teamID, firstPersonID, secondPersonID)
	if err != nil {
		logger.Log.WithError(err).Errorln("swap duties error")
		return nil, err
	}

	return view, nil
}

func (s *Service) Assign(ctx context.Context, teamID int64, personID int64) (*models.StatusView, error) {
	view, err := s.client.Assign(ctx, teamID, personID)
	if err != nil {
		logger.Log.WithError(err).Errorln("assign duty error")
		return nil, err
	}

	return view, nil
}

func (s *Service) GetCurrentDuty(ctx context.Context, teamID int64) (*models.DutyView, error) {
	view, err := s.client.GetCurrentDuty(ctx, teamID)
	if err != nil {
		logger.Log.WithError(err).Errorln("get current duty error")
		return nil, err
	}

	return view, nil
}

func (s *Service) GetDuties(
	ctx context.Context, teamID int64, count int64) ([]models.DutyView, error) {
	view, err := s.client.GetDuties(ctx, teamID, count)
	if err != nil {
		logger.Log.WithError(err).Errorln("get duties error")
		return nil, err
	}

	return view, nil
}
