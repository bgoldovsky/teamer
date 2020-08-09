package main

import (
	"github.com/bgoldovsky/dutyer/gateway-api/internal/cfg"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/teams"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/handlers"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	teamsRepo "github.com/bgoldovsky/dutyer/gateway-api/internal/repostiory/teams"
	teamsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/teams"
)

func main() {
	teamsHost := cfg.GetTeamsHost()
	secret := cfg.GetSecret()
	port := cfg.GetHTTPPort()
	redisAddress := cfg.GetRedisAddress()

	teamsClient, err := teams.NewClient(teamsHost)
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
