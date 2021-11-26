//serivodes fulcrum son bakanes aaaaa POR LA CONCHETUMAREEEE ahora deberia poderse vamo chupete suazo
package main

import (
	//"context"
	"context"
	"fmt"
	"io/ioutil"
	"log"

	//"net"
	"os"
	"strings"

	pb "github.com/fabiusinfo/StarWars/proto"
	//	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStarWarsServiceServer
}

func (s *server) SendInformationF(ctx context.Context, in *pb.SendRequest) (*pb.SendReply2, error) {

	return &pb.SendReply2{Message: "Fulcrum recibió tu información con éxito"}, nil
}

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

/*
func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {

//aqui implementar la escritura del archivo de texto
	//commando := GetMessage()
	comando := "Latierra4 Valpo 8"
	palabra := strings.Split(commando, " ")
	planeta := palabra[0]
	var path = "RP/" + planeta + ".txt"

		if delet == 1 {
		delet = 0
			nombreArchivo = path // El nombre o ruta absoluta del archivo
			err := osRemove(nombreArchivo)
			if err != nil {
				fmt.Printf("Error eliminando achivo: %v\n", err)
			} else {
				fmt.Println(" ")
			}
		}

	crarArchivo(path)

	// añadir al texto
	b, errtxt := ioutil.ReadFile(path)

	if errtxt != nil {
	log.Fatal(errtxt)
	}

	nmbre_ciudad := palabra[1]
	cantidad_soldados_rebeldes : palabra[2]

	b = append(b, []bye( nombre_planea + " " +nombre_ciudad +" " + cantidad_soldados_rebeldes +" \n")...)
	errtxt = ioutil.WrteFile(path, b, 0644)

	if errtxt != nil {
		log.Fatal(errtxt)
		}

	fm.Println("Se recibe... Player: " + in.GetPlayer() + " / Play:  " + in.GetPlay() + " / Stage: " + in.GetStage())
	return &pb.SndReply{Stage: "Amongus", Alive: true}, nil
}
*/
func main() {
	//nos convertios en servidor
	/*
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
	
	fmt.Println("Esperando un: oye!")
	fmt.Scanln(&X)
*/
	//aqui implementar la escritura del archivo de texto

	

	comando := "Latierra4 Valpo 8"
	palabra := strings.Split(comando, " ")
	nombre_planeta := palabra[0]
	nombre_ciudad := palabra[1]
	cantidad_soldados_rebeldes := palabra[2]
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

}
