package main

import (
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/cfg"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/clients/teams"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/handlers"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/logger"
)

func main() {
	peopleHost := cfg.GetPeopleHost()
	secret := cfg.GetSecret()
	port := cfg.GetHTTPPort()

	client, err := teams.NewClient(peopleHost)
	if err != nil {
		logger.Log.WithError(err).Fatal("can't connect team service service-teams")
	}

	app := handlers.New(client)
	app.Initialize(secret)
	app.Run(port)
}
