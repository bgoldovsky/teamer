package duties

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models/scheduler"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/publisher"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/duties"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/persons"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/teams"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/logger"
)

const (
	topicDuties       = "duties"
	eventDutyAssigned = "duty.assigned"
)

var (
	ErrTeamsNotFound = errors.New("teams not found")
)

type Service struct {
	personsRepo persons.Repository
	dutiesRepo  duties.Repository
	teamsRepo   teams.Repository
	publisher   publisher.Publisher
}

func New(
	personsRepo persons.Repository,
	dutiesRepo duties.Repository,
	teamsRepo teams.Repository,
	publisher publisher.Publisher,
) *Service {
	return &Service{
		personsRepo: personsRepo,
		dutiesRepo:  dutiesRepo,
		teamsRepo:   teamsRepo,
		publisher:   publisher,
	}
}

func (s *Service) GetCurrentDuty(ctx context.Context, teamID int64) (*models.Duty, error) {
	duty, err := s.dutiesRepo.Get(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("get current duty error: %w", err)
	}
	return duty, nil
}

func (s *Service) GetDuties(ctx context.Context, teamID int64, count int64, date time.Time) ([]models.Duty, error) {
	duty, err := s.dutiesRepo.Get(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("get current duty error: %w", err)
	}

	order := getOrder(duty)

	filter := &v1.PersonFilter{TeamIds: []int64{teamID}}
	personList, err := s.personsRepo.GetList(ctx, filter, 1000, 0, "asc", "id")
	if err != nil {
		return nil, fmt.Errorf("get persons error: %w", err)
	}

	dutyList, err := scheduler.Schedule(order, personList, date, count)
	if err != nil {
		return nil, fmt.Errorf("schedure duties error: %w", err)
	}

	return dutyList, nil
}

func (s *Service) Swap(ctx context.Context, request *v1.SwapRequest) error {
	err := s.dutiesRepo.Swap(ctx, request.TeamId, request.FirstPersonId, request.SecondPersonId)
	if err != nil {
		return fmt.Errorf("swap persons error: %w", err)
	}
	return nil
}

func (s *Service) Assign(ctx context.Context, teamID int64, personID int64, date time.Time) error {
	person, err := s.personsRepo.Get(ctx, personID)
	if err != nil {
		return fmt.Errorf("get person error: %w", err)
	}

	if person.TeamID != teamID {
		return errors.New("invalid team ID error")
	}

	duty := &models.Duty{
		TeamID:    person.TeamID,
		PersonID:  person.ID,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Slack:     person.Slack,
		Order:     person.DutyOrder,
		Month:     date.Month(),
		Day:       int64(date.Day()),
	}

	err = s.dutiesRepo.Save(ctx, duty)
	if err != nil {
		return fmt.Errorf("save duty error: %w", err)
	}

	err = s.publisher.Publish(eventDutyAssigned, teamID, topicDuties)
	if err != nil {
		return fmt.Errorf("publish %s to %s error: %w", eventDutyAssigned, topicDuties, err)
	}

	return nil
}

func (s *Service) AssignNextDuties(ctx context.Context, date time.Time) error {
	teamList, err := s.teamsRepo.GetList(ctx, nil, 1000, 0, "asc", "id")
	if err != nil {
		return fmt.Errorf("get team list error: %w", err)
	}

	if len(teamList) == 0 {
		return ErrTeamsNotFound
	}

	// Костыль. Пока вместо метрик используем логирование.
	for _, team := range teamList {
		err = s.assignNextDuty(ctx, team.ID, date)
		if err != nil {
			logger.Log.WithError(err).WithField("teamID", team.ID).Error("assign next duty error")
			continue
		}
		logger.Log.WithField("teamID", team.ID).Info("assign next duty success")
	}

	return nil
}

func (s *Service) assignNextDuty(ctx context.Context, teamID int64, date time.Time) error {
	prevDuty, err := s.dutiesRepo.Get(ctx, teamID)
	if err != nil {
		return fmt.Errorf("get current duty error: %w", err)
	}

	filter := &v1.PersonFilter{TeamIds: []int64{teamID}}
	personList, err := s.personsRepo.GetList(ctx, filter, 1000, 0, "asc", "id")
	if err != nil {
		return fmt.Errorf("get persons error: %w", err)
	}

	if isDutyAssigned(prevDuty, date) {
		return errors.New("duty is already assigned")
	}

	nextOrder := getOrder(prevDuty) + 1
	nextDuty, err := scheduler.Current(nextOrder, personList, date)
	if err != nil {
		return fmt.Errorf("schedule duties error: %w", err)
	}

	err = s.dutiesRepo.Save(ctx, nextDuty)
	if err != nil {
		return fmt.Errorf("save duty error: %w", err)
	}

	err = s.publisher.Publish(eventDutyAssigned, teamID, topicDuties)
	if err != nil {
		return fmt.Errorf("publish %s to %s error: %w", eventDutyAssigned, topicDuties, err)
	}

	return nil
}

func isDutyAssigned(duty *models.Duty, date time.Time) bool {
	if duty != nil && duty.Month == date.Month() && duty.Day == int64(date.Day()) {
		return true
	}
	return false
}

func getOrder(duty *models.Duty) int64 {
	if duty == nil {
		return -1
	}
	return duty.Order
}
