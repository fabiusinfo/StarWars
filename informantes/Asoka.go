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

func Interface(){
	var action, planet, city, value string
    fmt.Println("Bienvenda Ahsoka Tano al <Registro planetario>.")
    flag := true
    for flag {
        fmt.Println("Ingrese el número del comando a usar:")
        fmt.Println("-1- AddCity\n-2- UpdateName\n-3- UpdateNumber\n-4- DeleteCity\n-5- Salir")
        fmt.Scanln(&action)
        
        if action == "1" || action == "2"  || action == "3"{
            fmt.Println("Ingrese -Nombre planeta-")
            fmt.Scanln(&planet)
            
            fmt.Println("Ingrese -Nombre ciudad-")
            fmt.Scanln(&city)
            
            fmt.Println("Ingrese -Nuevo valor-")
            fmt.Scanln(&value)

            //mandarselo al Broker
            //action, planet, city, value
        } else if action == "4" {
            fmt.Println("Ingrese -Nombre planeta-")
            fmt.Scanln(&planet)
            
            fmt.Println("Ingrese -Nombre ciudad-")
            fmt.Scanln(&city)
            
            //mandarselo al Broker
            //action, planet, city
        } else if action == "5" {
            flag = false
            
            //finaliza el proceso del informante
        } else {
            fmt.Println("Ingrese un -comando válido-")
        }
    }
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
