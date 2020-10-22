package persons

import (
	"context"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/persons"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	personsRepo "github.com/bgoldovsky/dutyer/gateway-api/internal/repostiory/persons"
	"github.com/sirupsen/logrus"
)

type Service struct {
	client *persons.Client
	repo   personsRepo.Repository
}

func New(client *persons.Client, repo personsRepo.Repository) *Service {
	return &Service{
		client: client,
		repo:   repo,
	}
}

func (s *Service) GetPerson(ctx context.Context, personID int64) (*models.PersonView, error) {
	view, err := s.client.GetPerson(ctx, personID)
	if err != nil {
		logger.Log.WithError(err).Errorln("get person error")
		return nil, err
	}

	return view, nil
}

func (s *Service) GetPersons(ctx context.Context, teamID *int64) ([]models.PersonView, error) {
	if teamID == nil {
		return s.getAllPersons(ctx)
	}

	view, err := s.client.GetPersons(ctx, teamID)
	if err != nil {
		logger.Log.WithError(err).Errorln("get persons error")
		return nil, err
	}

	return view, nil
}

func (s *Service) getAllPersons(ctx context.Context) ([]models.PersonView, error) {
	if view := s.getRepo(); view != nil {
		return view, nil
	}

	view, err := s.client.GetPersons(ctx, nil)
	if err != nil {
		logger.Log.WithError(err).Errorln("get persons error")
		return nil, err
	}

	s.saveRepo(view)
	return view, nil
}

func (s *Service) AddPerson(ctx context.Context, form *models.PersonForm) (*models.StatusView, error) {
	status, err := s.client.AddPerson(ctx, form)
	if err != nil {
		logger.Log.WithField("form", form).WithError(err).Errorln("add person error")
		return nil, err
	}

	s.clearRepo()
	return status, nil
}

func (s *Service) UpdatePerson(ctx context.Context, personID int64, form *models.PersonForm) (*models.StatusView, error) {
	status, err := s.client.UpdatePerson(ctx, personID, form)
	if err != nil {
		logger.Log.
			WithFields(logrus.Fields{"form": form, "personID": personID}).
			WithError(err).Errorln("update person error")
		return nil, err
	}

	s.clearRepo()
	return status, nil
}

func (s *Service) RemovePerson(ctx context.Context, personID int64) (*models.StatusView, error) {
	status, err := s.client.RemovePerson(ctx, personID)
	if err != nil {
		logger.Log.WithField("personID", personID).WithError(err).Errorln("remove person error")
		return nil, err
	}

	s.clearRepo()
	return status, nil
}

func (s *Service) clearRepo() {
	err := s.repo.Clear()
	if err != nil {
		logger.Log.WithError(err).Error("clear persons repo error")
	}
}

func (s *Service) saveRepo(person []models.PersonView) {
	err := s.repo.Save(person)
	if err != nil {
		logger.Log.WithError(err).Error("save persons repo error")
	}
}

func (s *Service) getRepo() []models.PersonView {
	views, err := s.repo.Get()
	if err != nil {
		logger.Log.WithError(err).Error("get persons repo error")
		return nil
	}
	return views
}
