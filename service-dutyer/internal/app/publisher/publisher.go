//go:generate mockgen -destination publisher_mock/publisher_mock.go -source publisher.go

package publisher

import (
	"fmt"
	"time"

	dataBus "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/data-bus/v1"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/logger"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	_ "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka/librdkafka"
)

type Publisher interface {
	Publish(eventName string, entityID int64, topic string) error
}

type publisher struct {
	producer *kafka.Producer
}

func New(address string) (*publisher, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": address, "log.queue": "false"})
	if err != nil {
		return nil, fmt.Errorf("publisher create error: %w", err)
	}
	pub := &publisher{producer: producer}
	go pub.watchEvents()

	return pub, nil
}

func (p *publisher) watchEvents() {
	for e := range p.producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				logger.Log.WithField("partition", ev.TopicPartition).Info("delivery failed")
			} else {
				logger.Log.WithField("partition", ev.TopicPartition).Info("delivered message")
			}
		}
	}
}

func (p *publisher) Publish(eventName string, entityID int64, topic string) error {
	event, err := newEvent(eventName, entityID)
	if err != nil {
		return err
	}

	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          event,
	}, nil)

	if err != nil {
		return err
	}

	p.producer.Flush(15 * 1000)
	return nil
}

func newEvent(eventName string, entityID int64) ([]byte, error) {
	messageID, err := newMessageID()
	if err != nil {
		return nil, err
	}

	msg := &dataBus.EventMessage{
		MessageID: messageID,
		Data: &dataBus.EventData{
			Event:      eventName,
			EntityID:   entityID,
			OccurredOn: time.Now().UTC().Unix(),
		},
	}

	return proto.Marshal(msg)
}

func newMessageID() ([]byte, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	messageID := [16]byte(id)
	return messageID[:], nil
}
