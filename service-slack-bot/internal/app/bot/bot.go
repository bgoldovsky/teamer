package bot

import (
	"context"
	"fmt"

	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/logger"
	"github.com/slack-go/slack"
)

type SlackBot struct {
	slackClient *slack.Client
}

func NewSlackBot(slackToken string) *SlackBot {
	bot := &SlackBot{
		slackClient: slack.New(slackToken),
	}
	return bot
}

func (s *SlackBot) ChangeTopic(ctx context.Context, slack string, channel string) error {
	msg := fmt.Sprintf(template, slack)
	_, err := s.slackClient.SetTopicOfConversationContext(ctx, channel, msg)
	if err != nil {
		logger.Log.WithError(err).WithField("channel", channel).Error("set channel topic error")
		return fmt.Errorf("set channel topic error: %w", err)
	}
	return nil
}
