package main

import (
	//"context"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	//"math/rand"
	//"strconv"
	//"time"

	pb "github.com/fabiusinfo/StarWars/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStarWarsServiceServer
}

func (s *server) ConsultPlanet(ctx context.Context, in *pb.ConsultRequest) (*pb.ConsultReply, error) {

	// alo fulcrum paha toa la info
	direction := "10.6.43.44"
	conn, err := grpc.Dial(direction+":9000", grpc.WithInsecure())
	serviceSF := pb.NewStarWarsServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := serviceSF.ConsultPlanet(ctx, &pb.ConsultRequest{Message: "holi soy el broker"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// ahora esta respuesta se la mandamos a la leia
	return &pb.ConsultReply{Message: r.GetMessage()}, nil
}

func (s *server) SendInformationB(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {

	
	// conexion con informantes
	conn, err := grpc.Dial("10.6.43.43:8080", grpc.WithInsecure())
	
	if err != nil {
		panic("cannot connect with server " + err.Error())
	}
	
	serviceBroker := pb.NewStarWarsServiceClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := serviceBroker.SendInformationB(ctx, &pb.SendRequest{Message: "algo"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return &pb.SendReply{Ip: r.getIp() , Port: r.getPort()}, nil
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

/*

// Conexion con los servidores Fulcrum

rand.Seed(time.Now().UnixNano())
id := rand.Int63n(3)
if id == 0 {
	direction = "10.6.43.42" // maquina 2
} else if id == 1 {
	direction = "10.6.43.43" // maquina 3
} else {
	direction = "10.6.43.44" // maquina 4
}
conn, err := grpc.Dial(direction+":9000", grpc.WithInsecure())
serviceSF := pb.NewStarWarsServiceClient(conn)
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
_, err = serviceSF.SendPlays(ctx, &pb.SendRequest{Player: "jaja nose"})
if err != nil {
	log.Fatalf("could not greet: %v", err)
}


*/
