package main

import (
	//"context"
	//"fmt"
	//"log"
	//"math"
	//"math/rand"
	//"net"
	//"strconv"
	//"time"

	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/fabiusinfo/StarWars/proto"
	//amqp "github.com/rabbitmq/amqp091-go"
	//"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStarWarsServiceServer
}

func main() {
	action := "none"
	flag1 := false
	for !flag1 {
		fmt.Println("Escribe -consultar- para consultar un planeta ")
		fmt.Scanln(&action)
		if action == "consultar" {
			flag1 = true
		}
	}
	conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	servicePlayer := pb.NewSquidGameServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := servicePlayer.ConsultPlanet(ctx, &pb.JoinRequest{Message: ""})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}
