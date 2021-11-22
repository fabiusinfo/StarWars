package main

import (
	//"context"
	"context"
	"fmt"
	"log"
	"net"

	//"strconv"
	//"time"

	pb "github.com/fabiusinfo/StarWars/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStarWarsServiceServer
}

type PlayerStruct struct {
	id    string
	alive bool
	score int
}

var liderPlay int
var actualStage string
var actualRound int32
var started bool
var flaggy bool
var list_of_players []PlayerStruct

//listas stage 2
var group1 []PlayerStruct
var group2 []PlayerStruct
var groupaux []PlayerStruct

//listas stage 3
var group3 []PlayerStruct

//var players [16]string
var totalPlayers int

// Error para el Rabbit
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (s *server) ConsultPlanet(ctx context.Context, in *pb.ConsultRequest) (*pb.ConsultReply, error) {
	return &pb.ConsultReply{Message: "toma la info del planeta"}, nil
}

func (s *server) SendInformation(ctx context.Context, in *pb.SendRequest) (*pb.sendReply, error) {

	return &pb.SendReply{Message: "información recibida con éxito"}, nil
}

func main() {
	X := "none"
	go func() {
		// nos convertimos en servidor
		listner, err := net.Listen("tcp", ":8080")

		if err != nil {
			panic("cannot create tcp connection" + err.Error())
		}

		serv := grpc.NewServer()
		pb.RegisterStarWarsServiceServer(serv, &server{})

		//esto es lo que estaba al final, no sé donde ponerlo
		if err = serv.Serve(listner); err != nil {
			log.Printf("paso por el fallo")
			panic("cannot initialize the server" + err.Error())
		}

	}()
	fmt.Println("Esperando un: oye!")
	fmt.Scanln(&X)

}
