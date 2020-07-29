package main

import (
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/cfg"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/clients/teams"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/handlers"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/logger"
	teamsRepo "github.com/bgoldovsky/teamer-bot/gateway-api/internal/repostiory/teams"
	teamsSrv "github.com/bgoldovsky/teamer-bot/gateway-api/internal/services/teams"
)

func main() {
	peopleHost := cfg.GetPeopleHost()
	secret := cfg.GetSecret()
	port := cfg.GetHTTPPort()
	redisAddress := cfg.GetRedisAddress()

	teamsClient, err := teams.NewClient(peopleHost)
	if err != nil {
		logger.Log.WithError(err).Fatal("can't connect team service service-teams")
	}

	teamsRepository, err := teamsRepo.NewRepository(redisAddress)
	if err != nil {
		logger.Log.WithError(err).Fatal("can't connect redis")
	}

	teamsService := teamsSrv.New(teamsClient, teamsRepository)

	app := handlers.New(teamsService)
	app.Initialize(secret)
	app.Run(port)
}
