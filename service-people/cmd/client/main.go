package main

import (
	"context"
	"fmt"
	"log"
	"time"

	v1 "github.com/bgoldovsky/teamer-bot/service-people/internal/generated/rpc/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := v1.NewTeamsClient(conn)

	Get(client)
}

func Get(client v1.TeamsClient) {
	to, _ := ptypes.TimestampProto(time.Now())
	filter := &v1.TeamFilter{
		Ids:    []int64{101},
		DateTo: to,
		//DateFrom: from,
	}
	fmt.Println(filter)

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

func Update(client v1.TeamsClient) {
	_, err := client.UpdateTeam(context.Background(), &v1.UpdateTeamRequest{
		Team: &v1.Team{
			Id:          168,
			Name:        "Avito 123",
			Description: "Mem team",
			Slack:       "POP!!!",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Remove(client v1.TeamsClient) {
	_, err := client.RemoveTeam(context.Background(), &v1.RemoveTeamRequest{
		Id: 168,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Add(client v1.TeamsClient) {
	res, err := client.AddTeam(context.Background(), &v1.AddTeamRequest{Name: "Super Team!", Description: "Best team ever", Slack: "XYZ"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res, err)
}
