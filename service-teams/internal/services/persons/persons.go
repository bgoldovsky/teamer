package persons

import (
	"context"

	v1 "github.com/bgoldovsky/dutyer/service-teams/internal/generated/rpc/v1"
	"github.com/bgoldovsky/dutyer/service-teams/internal/repository/persons"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	repo persons.Repository
}

func New(repo persons.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddPerson(ctx context.Context, req *v1.AddPersonRequest) (*v1.AddPersonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPerson not implemented")
}
func (s *Service) UpdatePerson(ctx context.Context, req *v1.UpdatePersonRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (s *Service) RemoverPerson(ctx context.Context, req *v1.RemovePersonRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoverPerson not implemented")
}
func (s *Service) GetPersons(ctx context.Context, req *v1.GetPersonsRequest) (*v1.GetPersonsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersons not implemented")
}
