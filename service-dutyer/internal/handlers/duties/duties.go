package duties

import (
	"context"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/services/duties"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/logger"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	service *duties.Service
}

func New(service *duties.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetCurrentDuty(ctx context.Context, req *v1.GetCurrentDutyRequest) (*v1.GetCurrentDutyReply, error) {
	if req.TeamId == 0 {
		return nil, status.Error(codes.InvalidArgument, "team id not specified")
	}

	reply, err := h.service.GetCurrentDuty(ctx, req.TeamId)
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Errorln("get current duty error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.GetCurrentDutyReply{Duty: toReply(reply)}, nil
}

func (h *Handler) GetDuties(ctx context.Context, req *v1.GetDutiesRequest) (*v1.GetDutiesReply, error) {
	if req.TeamId == 0 {
		return nil, status.Error(codes.InvalidArgument, "team id not specified")
	}

	if req.Count == 0 {
		return nil, status.Error(codes.InvalidArgument, "count not specified")
	}

	reply, err := h.service.GetDuties(ctx, req.TeamId, req.Count, time.Now())
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Errorln("get duties error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.GetDutiesReply{Duties: toReplies(reply)}, nil
}

func (h *Handler) Assign(ctx context.Context, req *v1.AssignRequest) (*empty.Empty, error) {
	if req.TeamId == 0 {
		return nil, status.Error(codes.InvalidArgument, "team id not specified")
	}

	if req.PersonId == 0 {
		return nil, status.Error(codes.InvalidArgument, "person id not specified")
	}

	err := h.service.Assign(ctx, req.TeamId, req.PersonId, time.Now())
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Error("assign duty error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (h *Handler) Swap(ctx context.Context, req *v1.SwapRequest) (*empty.Empty, error) {
	if req.TeamId == 0 {
		return nil, status.Error(codes.InvalidArgument, "team id not specified")
	}

	if req.FirstPersonId == 0 {
		return nil, status.Error(codes.InvalidArgument, "first person id not specified")
	}

	if req.SecondPersonId == 0 {
		return nil, status.Error(codes.InvalidArgument, "second person id not specified")
	}

	err := h.service.Swap(ctx, req)
	if err != nil {
		logger.Log.WithField("req", req).WithError(err).Error("swap duties error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func toReply(duty *models.Duty) *v1.Duty {
	if duty == nil {
		return nil
	}

	return &v1.Duty{
		TeamId:    duty.TeamID,
		PersonId:  duty.PersonID,
		FirstName: duty.FirstName,
		LastName:  duty.LastName,
		Slack:     duty.Slack,
		Channel:   duty.Channel,
		DutyOrder: duty.Order,
		Month:     int64(duty.Month),
		Day:       duty.Day,
	}
}

func toReplies(duties []models.Duty) []*v1.Duty {
	if duties == nil {
		return []*v1.Duty{}
	}

	reply := make([]*v1.Duty, len(duties))
	for idx, duty := range duties {
		reply[idx] = toReply(&duty)
	}

	return reply
}
