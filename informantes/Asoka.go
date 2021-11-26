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

func Interface() []string {
	var action, command, planet, city, value string
	var info []string
    flag := true
    for flag {
        fmt.Println("Ingrese el número del comando a usar:")
        fmt.Println("-1- AddCity\n-2- UpdateName\n-3- UpdateNumber\n-4- DeleteCity")
        fmt.Scanln(&action)
        
        if action == "1" || action == "2"  || action == "3"{
            fmt.Println("Ingrese -Nombre planeta-")
            fmt.Scanln(&planet)
            
            fmt.Println("Ingrese -Nombre ciudad-")
            fmt.Scanln(&city)
            
            fmt.Println("Ingrese -Nuevo valor-") //si no quiere en el 1 agregar nuevo valor es 0
            fmt.Scanln(&value)

			if action == "1" {
				command = "AddCity"
			} else if action == "2" {
				command = "UpdateName"
			} else {
				command = "UpdateNumber"
			}

			info = append(info, command, planet,city,value)

        } else if action == "4" {
            fmt.Println("Ingrese -Nombre planeta-")
            fmt.Scanln(&planet)
            
            fmt.Println("Ingrese -Nombre ciudad-")
            fmt.Scanln(&city)

			value = "0"
			info = append(info, "DeleteCity", planet,city,value)

        } else {
            fmt.Println("Ingrese un -comando válido-")
        }
    }
    return info
}


func main() {
	fmt.Println("Bienvenda Ahsoka Tano al <Registro planetario>.")
	flag := true

	for flag {
		message := Interface()

		//envío al Broker
		conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

		if err != nil {
			panic("cannot connect with server " + err.Error())
		}
		serviceInformant := pb.NewStarWarsServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := serviceInformant.SendInformationB(ctx, &pb.SendRequest{Command: message[0], Planet: message[1], City: message[2], Value:message[3]})
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

		r2, err := serviceInformant2.SendInformationF(ctx2, &pb.SendRequest{Command: message[0], Planet: message[1], City: message[2], Value:message[3]})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r2.GetMessage())
	}

}
