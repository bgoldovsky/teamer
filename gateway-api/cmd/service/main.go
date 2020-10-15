package main

import (
	"github.com/bgoldovsky/dutyer/gateway-api/internal/cfg"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/persons"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/teams"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/handlers"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	personsRepo "github.com/bgoldovsky/dutyer/gateway-api/internal/repostiory/persons"
	teamsRepo "github.com/bgoldovsky/dutyer/gateway-api/internal/repostiory/teams"
	personsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/persons"
	teamsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/teams"
)

func main() {
	dutyerHost := cfg.GetTeamsHost()
	secret := cfg.GetSecret()
	port := cfg.GetHTTPPort()
	redisAddress := cfg.GetRedisAddress()

	// Clients
	teamsClient, err := teams.NewClient(dutyerHost)
	fatalOnError("can't connect team service service-dutyer", err)

	personsClient, err := persons.NewClient(dutyerHost)
	fatalOnError("can't connect persons service service-dutyer", err)

	// Repositories
	teamsRepository, err := teamsRepo.NewRepository(redisAddress)
	fatalOnError("can't connect redis", err)

	personsRepository, err := personsRepo.NewRepository(redisAddress)
	fatalOnError("can't connect redis", err)

	// Services
	teamsService := teamsSrv.New(teamsClient, teamsRepository)
	personsService := personsSrv.New(personsClient, personsRepository)

	// Handlers
	app := handlers.New(teamsService, personsService)
	app.Initialize(secret)
	app.Run(port)
}

func fatalOnError(msg string, err error) {
	if err != nil {
		logger.Log.WithError(err).Fatal(msg)
	}
}
