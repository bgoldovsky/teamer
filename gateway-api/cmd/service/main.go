package main

import (
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/cfg"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/clients/teams"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/handlers"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/logger"
	teamsRepo "github.com/bgoldovsky/teamer-bot/gateway-api/internal/repostiory/teams"
)

func main() {
	peopleHost := cfg.GetPeopleHost()
	secret := cfg.GetSecret()
	port := cfg.GetHTTPPort()
	redisAddress := cfg.GetRedisAddress()

	client, err := teams.NewClient(peopleHost)
	if err != nil {
		logger.Log.WithError(err).Fatal("can't connect team service service-teams")
	}

	repo, err := teamsRepo.NewRepository(redisAddress)
	if err != nil {
		logger.Log.WithError(err).Fatal("can't connect redis")
	}

	app := handlers.New(client, repo)
	app.Initialize(secret)
	app.Run(port)
}
