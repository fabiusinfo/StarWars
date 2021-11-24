package main

import (
	"context"
	"fmt"
	"log"

	//"math"
	//"math/rand"
	//"net"
	//"strconv"
	"time"

	pb "github.com/fabiusinfo/StarWars/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStarWarsServiceServer
}

func main() {
	action := "none"
	flag1 := false
	for !flag1 {
		fmt.Println("Para enviar informacion escriba -send- ")
		fmt.Scanln(&action)
		if action == "send" {
			flag1 = true
		}
	}

	//envío al Broker
	conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}
	serviceInformant := pb.NewStarWarsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := serviceInformant.SendInformationB(ctx, &pb.SendRequest{Message: ""})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetIp())
	log.Printf("Greeting: %s", r.GetPort())

	//envío al Flucrum

	conn2, err := grpc.Dial(r.GetIp()+":"+r.GetPort(), grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}
	serviceInformant2 := pb.NewStarWarsServiceClient(conn2)

	ctx2, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r2, err := serviceInformant2.SendInformationF(ctx2, &pb.SendRequest{Message: "actualiza esto"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r2.GetMessage())
}
