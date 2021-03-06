package persons

import (
	"context"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/services/persons"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/logger"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	service *persons.Service
}

func New(service *persons.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetPerson(ctx context.Context, req *v1.GetPersonRequest) (*v1.GetPersonReply, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "person id not specified")
	}

	reply, err := h.service.GetPerson(ctx, req)
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Errorln("get person error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return reply, nil
}

func (h *Handler) GetPersons(ctx context.Context, req *v1.GetPersonsRequest) (*v1.GetPersonsReply, error) {
	if req.Order != "id" && req.Order != "name" {
		return nil, status.Error(codes.InvalidArgument, "order must be id|name")
	}

	if req.Sort != "asc" && req.Sort != "desc" {
		return nil, status.Error(codes.InvalidArgument, "sort must be asc|desc")
	}

	reply, err := h.service.GetPersons(ctx, req)
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Errorln("get persons error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return reply, nil
}

func (h *Handler) AddPerson(ctx context.Context, req *v1.AddPersonRequest) (*v1.AddPersonReply, error) {
	if len(req.FirstName) == 0 {
		return nil, status.Error(codes.InvalidArgument, "person first name not specified")
	}

	if len(req.LastName) == 0 {
		return nil, status.Error(codes.InvalidArgument, "person last name not specified")
	}

	if len(req.Slack) == 0 {
		return nil, status.Error(codes.InvalidArgument, "person slack not specified")
	}

	if req.Role == v1.Role_NONE {
		return nil, status.Error(codes.InvalidArgument, "person role not specified")
	}

	reply, err := h.service.AddPerson(ctx, req)
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Errorln("add person error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return reply, nil
}

func (h *Handler) UpdatePerson(ctx context.Context, req *v1.UpdatePersonRequest) (*empty.Empty, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "person id not specified")
	}

	if len(req.FirstName) == 0 {
		return nil, status.Error(codes.InvalidArgument, "person first name not specified")
	}

	if len(req.LastName) == 0 {
		return nil, status.Error(codes.InvalidArgument, "person last name not specified")
	}

	if len(req.Slack) == 0 {
		return nil, status.Error(codes.InvalidArgument, "person slack not specified")
	}

	if req.Role == v1.Role_NONE {
		return nil, status.Error(codes.InvalidArgument, "person role not specified")
	}

	reply, err := h.service.UpdatePerson(ctx, req)
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Errorln("update person error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return reply, nil
}

func (h *Handler) RemovePerson(ctx context.Context, req *v1.RemovePersonRequest) (*empty.Empty, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "person id not specified")
	}

	reply, err := h.service.RemoverPerson(ctx, req)
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Errorln("remove person error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return reply, nil
}
