package main

import (
	"github.com/bgoldovsky/teamer/service-duty-bot/internal/cfg"
	dataBus "github.com/bgoldovsky/teamer/service-duty-bot/internal/generated/data-bus/v1"
	"github.com/bgoldovsky/teamer/service-duty-bot/internal/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	_ "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka/librdkafka"
)

func main() {
	kafkaAddress := cfg.GetKafkaAddress()
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaAddress,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		logger.Log.WithError(err).Infoln("error connect broker")
		return
	}

	err = c.SubscribeTopics([]string{"teams"}, nil)
	if err != nil {
		logger.Log.WithError(err).Infoln("error subscribe topic")
		return
	}

	for {
		msg, err := c.ReadMessage(-1)
		//_, _ = c.Commit()
		//c.Commit()
		//c.CommitMessage(msg)

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
