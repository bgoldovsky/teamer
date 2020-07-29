package main

import (
	"context"
	"log"
	"net"

	"github.com/bgoldovsky/teamer/service-people/internal/cfg"

	"github.com/bgoldovsky/teamer/service-people/internal/database"
	v1 "github.com/bgoldovsky/teamer/service-people/internal/generated/rpc/v1"
	"github.com/bgoldovsky/teamer/service-people/internal/handlers/persons"
	"github.com/bgoldovsky/teamer/service-people/internal/handlers/teams"
	"github.com/bgoldovsky/teamer/service-people/internal/interceptors"
	"github.com/bgoldovsky/teamer/service-people/internal/logger"
	teamsRepo "github.com/bgoldovsky/teamer/service-people/internal/repository/teams"
	personsSrv "github.com/bgoldovsky/teamer/service-people/internal/services/persons"
	teamsSrv "github.com/bgoldovsky/teamer/service-people/internal/services/teams"
	"google.golang.org/grpc"
)

func main() {
	// Repo
	connString := cfg.GetConnString()
	var db = database.NewDatabase(context.Background(), connString)
	teamsRepository := teamsRepo.NewRepository(db)

	// Services
	personsService := personsSrv.New()
	teamsService := teamsSrv.New(teamsRepository)

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
