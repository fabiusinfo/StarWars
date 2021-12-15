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
	"google.golang.org/grpc"
	//amqp "github.com/rabbitmq/amqp091-go"
	//"google.golang.org/grpc"
)

type ReadStruct struct {
	command     string
	planet      string
	city        string
	Rebelds     string
	VectorClock string
	//ip          string
}

var MonotonicReads []ReadStruct

type server struct {
	pb.UnimplementedStarWarsServiceServer
}

func Interface() []string {
	var action, command, planet, city string
	var info []string
	flag := true

	for flag {
		fmt.Println("Ingrese el número del comando a usar:")
		fmt.Println("-1- GetNumberRebelds")
		fmt.Scanln(&action)

		if action == "1" {
			fmt.Println("Ingrese -Nombre planeta-")
			fmt.Scanln(&planet)

			fmt.Println("Ingrese -Nombre ciudad-")
			fmt.Scanln(&city)

			command = "GetNumberRebelds"
			info = append(info, command, planet, city)
			flag = false

		} else {
			fmt.Println("Ingrese un -comando válido-")
		}
	}
	return info
}

func main() {
	fmt.Println("Bienvenida Princesa Leia al <Registro planetario>.")
	flag := true

	for flag {
		message := Interface()

		//envío al Broker
		conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

		if err != nil {
			panic("cannot connect with server " + err.Error())
		}

		serviceLeya := pb.NewStarWarsServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := serviceLeya.ConsultPlanet(ctx, &pb.ConsultRequest{Command: message[0], Planet: message[1], City: message[2]})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Println("\nRebeldes: " + r.GetRebelds() + "\nReloj: " + r.GetClock()+"\n")

		MonotonicReads = append(MonotonicReads, ReadStruct{message[0], message[1], message[2], r.GetRebelds(), r.GetClock() /*, r.GetIp()*/})
		//fmt.Println(MonotonicReads)
	}
}
