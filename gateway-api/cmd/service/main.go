package main

import (
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/cfg"
	v1 "github.com/bgoldovsky/teamer-bot/gateway-api/internal/generated/clients/people/v1"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/handlers"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/interceptors"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/logger"
	"google.golang.org/grpc"
)

func main() {
	peopleHost := cfg.GetPeopleHost()
	secret := cfg.GetSecret()
	port := cfg.GetHTTPPort()

	conn, err := grpc.Dial(peopleHost, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptors.LoggingInterceptor))
	if err != nil {
		logger.Log.WithError(err).Fatal("can't connect service-people")
	}
	client := v1.NewTeamsClient(conn)

	app := handlers.New(client)
	app.Initialize(secret)
	app.Run(port)
}
