package main

import (
	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/cfg"
	dataBus "github.com/bgoldovsky/dutyer/service-slack-bot/internal/generated/data-bus/v1"
	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	_ "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka/librdkafka"
)

const (
	topicDuties = "duties"
)

func main() {
	kafkaAddress := cfg.GetKafkaAddress()
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaAddress,
		"group.id":          "xxx",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		logger.Log.WithError(err).Infoln("error connect broker")
		return
	}

	err = c.SubscribeTopics([]string{topicDuties}, nil)
	fatalOnError("error subscribe topic", err)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			logger.Log.WithError(err).Error("read message error")
		}

		_, _ = c.CommitMessage(msg)

		if err == nil {
			var event dataBus.EventMessage
			err = proto.Unmarshal(msg.Value, &event)
			if err != nil {
				logger.Log.
					WithField("event", msg.Value).
					WithField("topic", msg.TopicPartition).
					Infoln("error unmarshal message")
				continue
			}
			logger.Log.
				WithField("event", event).
				WithField("topic", msg.TopicPartition).
				Infoln("message consumed")
		} else {
			logger.Log.
				WithError(err).
				Infoln("consumer error")
		}
	}
}

func fatalOnError(msg string, err error) {
	if err != nil {
		logger.Log.WithError(err).Fatal(msg)
	}
}
