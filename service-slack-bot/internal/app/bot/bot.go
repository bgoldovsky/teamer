//go:generate mockgen -destination bot_mock/bot_mock.go -source bot.go

package bot

import (
	"context"
	"fmt"

	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/logger"
	"github.com/slack-go/slack"
)

type SlackBot interface {
	ChangeTopic(ctx context.Context, slack string, channel string) error
}

type slackBot struct {
	slackClient *slack.Client
}

func NewSlackBot(slackToken string) *slackBot {
	bot := &slackBot{
		slackClient: slack.New(slackToken),
	}
	return bot
}

func (s *slackBot) ChangeTopic(ctx context.Context, slack string, channel string) error {
	msg := fmt.Sprintf(template, slack)
	_, err := s.slackClient.SetTopicOfConversationContext(ctx, channel, msg)
	if err != nil {
		logger.Log.WithError(err).WithField("channel", channel).Error("set channel topic error")
		return fmt.Errorf("set channel topic error: %w", err)
	}
	return nil
}
