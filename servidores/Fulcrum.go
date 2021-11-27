//serivodes fulcrum son bakanes aaaaa POR LA CONCHETUMAREEEE ahora deberia poderse vamo chupete suazo
package main

import (
	//"context"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"

	//"net"
	"os"

	pb "github.com/fabiusinfo/StarWars/proto"
	"google.golang.org/grpc"
	//	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStarWarsServiceServer
}

//Global Variables
var VectorClock []int

/*func (s *server) SendInformationF(ctx context.Context, in *pb.SendRequest) (*pb.SendReply2, error) {

	return &pb.SendReply2{Message: "Fulcrum recibió tu información con éxito"}, nil
}*/

func (s *server) ConsultPlanet(ctx context.Context, in *pb.ConsultRequest) (*pb.ConsultReply, error) {

	return &pb.ConsultReply{Message: "toma la info del planeta"}, nil
}

// Crear archivo

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func crearArchivo(path string) {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
}

//var delet int = 1

func (s *server) SendInformationF(ctx context.Context, in *pb.SendRequest2) (*pb.SendReply2, error) {

	//aqui implementar la escritura del archivo de texto
	command := in.GetCommand()
	planet := in.GetPlanet()
	city := in.GetCity()
	value := in.GetValue()
	var path = "servidores/RP/" + planet + ".txt"
	fmt.Println("Comando recibido: " + command + " " + planet + " " + city + " " + value)

	crearArchivo(path)

	if command == "AddCity" {
		// añadir al texto
		b, errtxt := ioutil.ReadFile(path)

		if errtxt != nil {
			log.Fatal(errtxt)
		}

		b = append(b, []byte(planet+" "+city+" "+value+" \n")...)
		errtxt = ioutil.WriteFile(path, b, 0644)

		if errtxt != nil {
			log.Fatal(errtxt)
		}
	} else if command == "UpdateName" {

		input, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalln(err)
		}

		lines := strings.Split(string(input), "\n")

		for i, line := range lines {
			if strings.Contains(line, city) {
				splitLine := strings.Split(string(line), " ")
				soldiers := splitLine[2]
				lines[i] = planet + " " + value + " " + soldiers
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(path, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}

	} else if command == "UpdateNumber" {
		input, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalln(err)
		}

		lines := strings.Split(string(input), "\n")

		for i, line := range lines {
			if strings.Contains(line, city) {
				lines[i] = planet + " " + city + " " + value
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(path, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		fmt.Println("hi")
	}
	if in.GetFulcrum() == "1" {
		VectorClock[0] += 1
	} else if in.GetFulcrum() == "2" {
		VectorClock[1] += 1
	} else {
		VectorClock[2] += 1
	}

	return &pb.SendReply2{Message: strconv.Itoa(VectorClock[0]) + " " + strconv.Itoa(VectorClock[1]) + " " + strconv.Itoa(VectorClock[2])}, nil
}

func main() {
	//nos convertios en servidor
	//VectorClock := [3]int{0, 0, 0} //{f1-42, f2-43, f3-44}
	VectorClock = append(VectorClock, 0, 0, 0)

	X := "none"
	go func() {
		listner, err := net.Listen("tcp", ":9000")

		if err != nil {
			panic("cannot connect with server " + err.Error())
		}

		serv := grpc.NewServer()
		pb.RegisterStarWarsServiceServer(serv, &server{})
		if err = serv.Serve(listner); err != nil {
			panic("cannot initialize the server" + err.Error())

		}

	}()

	fmt.Println("<Servidor Fulcrum habilitado>")
	fmt.Scanln(&X)

	//aqui implementar la escritura del archivo de texto

	/*
		comando := "Latierra4 Valpo 8"
		palabra := strings.Split(comando, " ")
		nombre_planeta := in.GetPlanet()
		nombre_ciudad := in.GetCity()
		cantidad_soldados_rebeldes := in.GetValue()
		var path = "servidores/RP/"+nombre_planeta + ".txt"
		fmt.Println(nombre_planeta)
		fmt.Println(nombre_ciudad)
		fmt.Println(cantidad_soldados_rebeldes)

		crearArchivo(path)

		// añadir al texto
		b, errtxt := ioutil.ReadFile(path)

		if errtxt != nil {
			log.Fatal(errtxt)
		}

		b = append(b, []byte(nombre_ciudad+" "+cantidad_soldados_rebeldes+" \n")...)
		errtxt = ioutil.WriteFile(path, b, 0644)

		if errtxt != nil {
			log.Fatal(errtxt)
		}
	*/
}
