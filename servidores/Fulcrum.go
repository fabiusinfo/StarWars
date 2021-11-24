//serivodes fulcrum son bakanes aaaaa POR LA CONCHETUMAREEEE ahora deberia poderse vamo chupete suazo
package main

import (
	//"context"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	pb "github.com/fabiusinfo/StarWars/proto"
	"google.golang.org/grpc"
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

/*var delet int = 1

func (s *server) SedPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {

//aqui implementar la escritura del archivo de texto
	plneta := "tatuin"
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
	nombre_planeta := Tatooine"
	nmbre_ciudad := "Mos_Eisley"
	cantidad_soldados_rebeldes : "5"

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

	//aqui implementar la escritura del archivo de texto
	nombre_planeta := "Tatooine"
	nombre_ciudad := "mos_Eisley"
	cantidad_soldados_rebeldes := "5"
	var path = "RP/" + nombre_planeta + ".txt"
	/*
		if delet == 1 {
		elet = 0
			nombreArchivo = ath // El nombre o ruta absoluta del archivo
			err := osRemove(ombreArchivo)
		if err != nil {
			fmt.Printf("Error eliminando achivo: %v\n", err)
		} else {
				fmt.Println(" ")
			}
		}
	*/

	crearArchivo(path)

	// añadir al texto
	b, errtxt := ioutil.ReadFile(path)

	if errtxt != nil {
		log.Fatal(errtxt)
	}

	b = append(b, []byte(nombre_planeta+" "+nombre_ciudad+" "+cantidad_soldados_rebeldes+" \n")...)
	errtxt = ioutil.WriteFile(path, b, 0644)

	if errtxt != nil {
		log.Fatal(errtxt)
	}

}
