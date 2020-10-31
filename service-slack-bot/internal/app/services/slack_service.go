package services

import (
	"context"
	"fmt"

	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/app/bot"
	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/clients/duties"
)

type SlackService struct {
	bot    bot.SlackBot
	client duties.Client
}

func New(bot bot.SlackBot, client duties.Client) *SlackService {
	return &SlackService{
		bot:    bot,
		client: client,
	}
}

func (s *SlackService) ChangeChannelTopic(ctx context.Context, teamID int64) error {
	duty, err := s.client.GetCurrentDuty(ctx, teamID)
	if err != nil {
		return fmt.Errorf("get duty error: %w", err)
	}

	err = s.bot.ChangeTopic(ctx, duty.Slack, duty.Channel)
	if err != nil {
		return fmt.Errorf("slack notification error: %w", err)
	}

	return nil
}
