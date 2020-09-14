package main

import (
	"context"
	"log"
	"net"

	"github.com/bgoldovsky/dutyer/service-teams/internal/app/handlers/persons"
	"github.com/bgoldovsky/dutyer/service-teams/internal/app/handlers/teams"
	"github.com/bgoldovsky/dutyer/service-teams/internal/app/publisher"
	personsRepo "github.com/bgoldovsky/dutyer/service-teams/internal/app/repository/persons"
	teamsRepo "github.com/bgoldovsky/dutyer/service-teams/internal/app/repository/teams"
	personsSrv "github.com/bgoldovsky/dutyer/service-teams/internal/app/services/persons"
	teamsSrv "github.com/bgoldovsky/dutyer/service-teams/internal/app/services/teams"
	"github.com/bgoldovsky/dutyer/service-teams/internal/cfg"
	"github.com/bgoldovsky/dutyer/service-teams/internal/database"
	v1 "github.com/bgoldovsky/dutyer/service-teams/internal/generated/rpc/v1"
	"github.com/bgoldovsky/dutyer/service-teams/internal/interceptors"
	"github.com/bgoldovsky/dutyer/service-teams/internal/logger"
	"google.golang.org/grpc"
)

func main() {
	// Repo
	connString := cfg.GetConnString()
	var db = database.NewDatabase(context.Background(), connString)
	teamsRepository := teamsRepo.NewRepository(db)
	personsRepository := personsRepo.NewRepository(db)

	// Publisher
	kafkaAddress := cfg.GetKafkaAddress()
	pub, err := publisher.New(kafkaAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Services
	personsService := personsSrv.New(personsRepository, pub)
	teamsService := teamsSrv.New(teamsRepository, pub)

	// Handlers
	personsHandler := persons.New(personsService)
	teamsHandler := teams.New(teamsService)

	// Registration
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptors.LoggingInterceptor))
	v1.RegisterPersonsServer(server, personsHandler)
	v1.RegisterTeamsServer(server, teamsHandler)

	// Start
	port := cfg.GetGRPCPort()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	logger.Log.WithField("port", port).Infoln("gRPC server starts..")
	log.Fatal(server.Serve(listener))
}
