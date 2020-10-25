package cron

import (
	"context"
	"time"

	dutiesSrv "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/services/duties"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/logger"
	"github.com/robfig/cron"
)

func Start(expression string, service *dutiesSrv.Service) {
	c := cron.New()

	err := c.AddFunc("0 0/10 * * * ?", func() {
		err := service.AssignNextDuties(context.Background(), time.Now())
		if err != nil {
			logger.Log.WithError(err).Error("cron schedule next duties error")
			return
		}
	})

	fatalOnError("start cron error", err)
	c.Start()
}

func fatalOnError(msg string, err error) {
	if err != nil {
		logger.Log.WithError(err).Fatal(msg)
	}
}
