package main

import (
	"context"
	"log"
	"net"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/cron"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/publisher"
	dutiesRepo "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/duties"
	personsRepo "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/persons"
	teamsRepo "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/teams"
	dutiesSrv "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/services/duties"
	personsSrv "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/services/persons"
	teamsSrv "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/services/teams"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/cfg"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/database"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/handlers/duties"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/handlers/persons"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/handlers/teams"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/interceptors"
	"github.com/bgoldovsky/dutyer/service-dutyer/internal/logger"
	"google.golang.org/grpc"
)

func main() {
	// Repo
	connString := cfg.GetConnString()
	var db = database.NewDatabase(context.Background(), connString)
	teamsRepository := teamsRepo.NewRepository(db)
	personsRepository := personsRepo.NewRepository(db)
	dutiesRepository := dutiesRepo.NewRepository(db)

	// Publisher
	kafkaAddress := cfg.GetKafkaAddress()
	pub, err := publisher.New(kafkaAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Services
	personsService := personsSrv.New(personsRepository)
	teamsService := teamsSrv.New(teamsRepository)
	dutyService := dutiesSrv.New(personsRepository, dutiesRepository, teamsRepository, pub)

	// Cron
	cron.Start(cfg.GetCron(), dutyService)

	// Handlers
	personsHandler := persons.New(personsService)
	teamsHandler := teams.New(teamsService)
	dutiesHandler := duties.New(dutyService)

	// Registration
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptors.LoggingInterceptor))
	v1.RegisterPersonsServer(server, personsHandler)
	v1.RegisterTeamsServer(server, teamsHandler)
	v1.RegisterDutiesServer(server, dutiesHandler)

	// Start
	port := cfg.GetGRPCPort()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	logger.Log.WithField("port", port).Infoln("gRPC server starts..")
	log.Fatal(server.Serve(listener))
}
