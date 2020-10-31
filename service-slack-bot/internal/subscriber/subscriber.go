package subscriber

import (
	"context"
	"fmt"

	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/app/services"
	dataBus "github.com/bgoldovsky/dutyer/service-slack-bot/internal/generated/data-bus/v1"
	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	_ "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka/librdkafka"
)

const (
	topicDuties = "duties"
)

type Subscriber struct {
	consumer *kafka.Consumer
	service  *services.SlackService
}

func NewSubscriber(consumer *kafka.Consumer, service *services.SlackService) (*Subscriber, error) {
	err := consumer.SubscribeTopics([]string{topicDuties}, nil)
	if err != nil {
		return nil, fmt.Errorf("subscribe topic error: %w", err)
	}

	return &Subscriber{
		consumer: consumer,
		service:  service,
	}, nil
}

func (s *Subscriber) Consume() {
	for {
		ctx := context.Background()

		msg, err := s.consumer.ReadMessage(-1)
		if err != nil {
			logger.Log.WithError(err).Error("read message error")
			continue
		}

		var event dataBus.EventMessage
		err = proto.Unmarshal(msg.Value, &event)
		if err != nil {
			logger.Log.WithError(err).Error("unmarshal message error")
			continue
		}

		err = s.service.ChangeChannelTopic(ctx, event.Data.EntityID)
		if err != nil {
			logger.Log.WithError(err).Error("change channel topic error")
			continue
		}

		_, err = s.consumer.CommitMessage(msg)
		if err != nil {
			logger.Log.WithError(err).Error("commit message error")
		}
	}
}
