package services

import (
	"context"
	"errors"
	"testing"

	mockBot "github.com/bgoldovsky/dutyer/service-slack-bot/internal/app/bot/bot_mock"
	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/app/models"
	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/clients/duties"
	"github.com/golang/mock/gomock"
)

var (
	duty = &models.Duty{
		TeamID:    777,
		PersonID:  888,
		FirstName: "Boris",
		LastName:  "Goldovsky",
		Slack:     "QWERTY",
		Channel:   "YTREWQ",
		Month:     12,
		Day:       31,
	}
)

func TestSlackService_ChangeChannelTopic_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	bot := mockBot.NewMockSlackBot(ctrl)
	bot.EXPECT().
		ChangeTopic(ctx, gomock.Any(), gomock.Any()).
		Return(nil)

	client := duties.NewMockClient(ctrl)
	client.EXPECT().
		GetCurrentDuty(ctx, duty.TeamID).
		Return(duty, nil)

	service := New(bot, client)

	act := service.ChangeChannelTopic(ctx, duty.TeamID)

	if act != nil {
		t.Errorf("expected no error, got: %v", act)
	}
}

func TestSlackService_ChangeChannelTopic_SlackError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	srvErr := errors.New("my expected error")
	expErr := "slack notification error: my expected error"

	client := duties.NewMockClient(ctrl)
	client.EXPECT().
		GetCurrentDuty(ctx, duty.TeamID).
		Return(duty, nil)

	bot := mockBot.NewMockSlackBot(ctrl)
	bot.EXPECT().
		ChangeTopic(ctx, gomock.Any(), gomock.Any()).
		Return(srvErr)

	service := New(bot, client)

	act := service.ChangeChannelTopic(ctx, duty.TeamID)

	if act.Error() != expErr {
		t.Errorf("expected: %v, got: %v", expErr, act)
	}
}

func TestSlackService_ChangeChannelTopic_ClientError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	srvErr := errors.New("my expected error")
	expErr := "get duty error: my expected error"

	client := duties.NewMockClient(ctrl)
	client.EXPECT().
		GetCurrentDuty(ctx, duty.TeamID).
		Return(nil, srvErr)

	service := New(nil, client)

	act := service.ChangeChannelTopic(ctx, duty.TeamID)

	if act.Error() != expErr {
		t.Errorf("expected: %v, got: %v", expErr, act)
	}
}
