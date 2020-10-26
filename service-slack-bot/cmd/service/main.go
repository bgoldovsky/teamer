package main

import (
	slackBot "github.com/bgoldovsky/dutyer/service-slack-bot/internal/app/bot"
	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/cfg"
	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/clients/duties"
	"github.com/bgoldovsky/dutyer/service-slack-bot/internal/logger"
	sub "github.com/bgoldovsky/dutyer/service-slack-bot/internal/subscriber"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Clients
	dutiesClient, err := duties.NewClient(cfg.GetDutyerHost())

	// Slack bot
	bot := slackBot.NewSlackBot(cfg.GetSlackToken())

	// Kafka
	address := cfg.GetKafkaAddress()
	consumer, err := kafka.NewConsumer(
		&kafka.ConfigMap{"bootstrap.servers": address, "group.id": "xxx", "auto.offset.reset": "earliest"},
	)
	defer consumer.Close()
	fatalOnError("create consumer error", err)

	subscriber, err := sub.NewSubscriber(consumer, bot, dutiesClient)
	fatalOnError("error connect broker", err)

	// Start consume events
	subscriber.Consume()
}

func fatalOnError(msg string, err error) {
	if err != nil {
		logger.Log.WithError(err).Fatal(msg)
	}
}
