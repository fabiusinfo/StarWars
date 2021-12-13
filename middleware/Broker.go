package main

import (
	//"context"
	"context"
	"fmt"
	"log"
	"math/rand"
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
	direction := ""
	rebels := ""
	clock := ""
	for i := 0; i < 3; i++ {
		if i == 0 {
			direction = "10.6.43.44"
		} else if i == 1 {
			direction = "10.6.43.42"
		} else {
			direction = "10.6.43.43"
		}
		conn, err := grpc.Dial(direction+":9000", grpc.WithInsecure())
		serviceSF := pb.NewStarWarsServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := serviceSF.ConsultPlanet(ctx, &pb.ConsultRequest{Command: in.GetCommand(), Planet: in.GetPlanet(), City: in.GetCity()})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		if r.GetRebelds() != "none" {
			rebels = r.GetRebelds()
			clock = r.GetClock()
			break
		}
	}
	// alo fulcrum paha toa la info

	// ahora esta respuesta se la mandamos a la leia
	return &pb.ConsultReply{Rebelds: rebels, Clock: clock}, nil
}

func (s *server) SendInformationB(ctx context.Context, in *pb.SendRequestB) (*pb.SendReplyB, error) {

	var direction string
	fulcrum1 := "10.6.43.42"
	fulcrum2 := "10.6.43.43"
	fulcrum3 := "10.6.43.44"

	rand.Seed(time.Now().UnixNano())
	id := rand.Int63n(3)
	if id == 0 {
		direction = fulcrum1 // maquina 2
	} else if id == 1 {
		direction = fulcrum2 // maquina 3
	} else {
		direction = fulcrum3 // maquina 4
	}
	return &pb.SendReplyB{Ip: direction, Port: "9000"}, nil
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
	fmt.Println("-ENTER- para enviar IP's a Fulcrums")
	fmt.Scanln(&X)
	//envío ips a servidores fulcrums
	ip1 := ""
	ip2 := ""
	ip3 := ""
	for i := 0; i < 3; i++ {
		if i == 0 {
			ip1 = "10.6.43.42"
			ip2 = "10.6.43.43"
			ip3 = "10.6.43.44"
		} else if i == 1 {
			ip1 = "10.6.43.43"
			ip2 = "10.6.43.42"
			ip3 = "10.6.43.44"
		} else {
			ip1 = "10.6.43.44"
			ip2 = "10.6.43.42"
			ip3 = "10.6.43.43"
		}

		conn, err := grpc.Dial(ip1+":9000", grpc.WithInsecure())

		if err != nil {
			panic("cannot connect with server " + err.Error())
		}
		serviceInformant := pb.NewStarWarsServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := serviceInformant.Identify(ctx, &pb.SendIp{Ip: ip1, Ip1: ip2, Ip2: ip3})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Println(r.GetMessage())

	}

	fmt.Println("<Servidor Broker habilitado>")
	fmt.Scanln(&X)

}
