package main

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/bgoldovsky/teamer-bot/service-people/internal/generated/rpc/v1"
	"google.golang.org/grpc"
)

// Test gRPC client
func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := v1.NewTeamsClient(conn)

	GetTeam(client)
}

func GetTeam(client v1.TeamsClient) {
	teams, err := client.GetTeams(context.Background(), &v1.GetTeamsRequest{
		Filter: nil,
		Limit:  100,
		Offset: 0,
		Order:  "id",
		Sort:   "desc",
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, t := range teams.Teams {
		log.Println(*t)
	}
}

func UpdateTeam(client v1.TeamsClient) {
	_, err := client.UpdateTeam(context.Background(), &v1.UpdateTeamRequest{
		Team: &v1.Team{
			Id:          168,
			Name:        "Dream 123",
			Description: "Dream team v.123",
			Slack:       "QWERTY",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func RemoveTeam(client v1.TeamsClient) {
	_, err := client.RemoveTeam(context.Background(), &v1.RemoveTeamRequest{
		Id: 168,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func AddTeam(client v1.TeamsClient) {
	res, err := client.AddTeam(
		context.Background(),
		&v1.AddTeamRequest{
			Name:        "Super Team!",
			Description: "Best team ever",
			Slack:       "XYZ"},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res, err)
}
