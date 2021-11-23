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

func (s *server) ConsultPlanet(ctx context.Context, in *pb.ConsultRequest) (*pb.ConsultReply, error) {
	return &pb.ConsultReply{Message: "toma la info del planeta"}, nil
}

func (s *server) SendInformationB(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {

	return &pb.SendReply{Ip: "xx.xx.xx.xx", Port: "xxxx"}, nil
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
