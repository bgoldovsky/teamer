package main

import (
	"github.com/bgoldovsky/dutyer/gateway-api/internal/cfg"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/duties"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/persons"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/teams"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/handlers"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	personsRepo "github.com/bgoldovsky/dutyer/gateway-api/internal/repostiory/persons"
	teamsRepo "github.com/bgoldovsky/dutyer/gateway-api/internal/repostiory/teams"
	dutiesSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/duties"
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

	dutiesClient, err := duties.NewClient(dutyerHost)
	fatalOnError("can't connect duties service service-dutyer", err)

	// Repositories
	teamsRepository, err := teamsRepo.NewRepository(redisAddress)
	fatalOnError("can't connect redis", err)

	personsRepository, err := personsRepo.NewRepository(redisAddress)
	fatalOnError("can't connect redis", err)

	// Services
	teamsService := teamsSrv.New(teamsClient, teamsRepository)
	personsService := personsSrv.New(personsClient, personsRepository)
	dutiesService := dutiesSrv.New(dutiesClient)

	// Handlers
	h := handlers.NewHandler(teamsService, personsService, dutiesService, secret)
	h.Run(port)
}

func fatalOnError(msg string, err error) {
	if err != nil {
		logger.Log.WithError(err).Fatal(msg)
	}
}
