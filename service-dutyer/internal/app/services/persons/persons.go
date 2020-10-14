package persons

import (
	"context"
	"fmt"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/persons"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/services/persons/converter"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Service struct {
	repo persons.Repository
}

func New(repo persons.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetPerson(ctx context.Context, req *v1.GetPersonRequest) (*v1.GetPersonReply, error) {
	person, err := s.repo.Get(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("get person error: %w", err)
	}
	return &v1.GetPersonReply{Person: converter.ToDTO(person)}, nil
}

func (s *Service) GetPersons(ctx context.Context, req *v1.GetPersonsRequest) (*v1.GetPersonsReply, error) {
	personModels, err := s.repo.GetList(ctx, req.Filter, uint(req.Limit), uint(req.Offset), req.Sort, req.Order)

	if err != nil {
		return nil, fmt.Errorf("get persons error: %w", err)
	}

	if len(personModels) == 0 {
		return &v1.GetPersonsReply{Persons: []*v1.Person{}}, nil
	}

	personsProto := make([]*v1.Person, len(personModels))
	for idx, person := range personModels {
		personsProto[idx] = converter.ToDTO(&person)
	}

	return &v1.GetPersonsReply{Persons: personsProto}, nil
}

func (s *Service) AddPerson(ctx context.Context, req *v1.AddPersonRequest) (*v1.AddPersonReply, error) {
	model := &models.Person{
		TeamID:    req.TeamId,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Birthday:  toTime(req.Birthday),
		Slack:     req.Slack,
		Role:      models.Role(req.Role),
		IsActive:  req.IsActive,
	}

	if req.MiddleName != nil {
		middleName := req.MiddleName.GetValue()
		model.MiddleName = &middleName
	}

	if req.Email != nil {
		email := req.Email.GetValue()
		model.Email = &email
	}

	if req.Phone != nil {
		phone := req.Phone.GetValue()
		model.Email = &phone
	}

	model, err := s.repo.Save(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("add person error: %w", err)
	}

	return &v1.AddPersonReply{Id: model.ID}, nil
}

func (s *Service) UpdatePerson(ctx context.Context, req *v1.UpdatePersonRequest) (*empty.Empty, error) {
	person := converter.ToModel(req)

	_, err := s.repo.Update(ctx, person)
	if err != nil {
		return nil, fmt.Errorf("update person error: %w", err)
	}

	return &empty.Empty{}, nil
}

func (s *Service) RemoverPerson(ctx context.Context, req *v1.RemovePersonRequest) (*empty.Empty, error) {
	_, err := s.repo.Remove(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("remove person error: %w", err)
	}
	return &empty.Empty{}, nil
}

func toTime(stamp *timestamp.Timestamp) *time.Time {
	t, _ := ptypes.Timestamp(stamp)
	return &t
}
