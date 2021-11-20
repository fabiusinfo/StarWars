//serivodes fulcrum son bakanes aaaaa POR LA CONCHETUMAREEEE ahora deberia poderse vamo chupete suazo
package main

import (
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
	
	/*
	//aqui implementar la escritura del archivo de texto
	planeta := "tatuin"
	var path = "RP/" + planeta + ".txt"
	
		if delet == 1 {
			delet = 0
			nombreArchivo := path // El nombre o ruta absoluta del archivo
			err := os.Remove(nombreArchivo)
			if err != nil {
				fmt.Printf("Error eliminando archivo: %v\n", err)
			} else {
				fmt.Println(" ")
			}
		}
	
	crearArchivo(path)

	// añadir al texto
	b, errtxt := ioutil.ReadFile(path)

	if errtxt != nil {
		log.Fatal(errtxt)
	}
	nombre_planeta := "Tatooine"
	nombre_ciudad := "Mos_Eisley"
	cantidad_soldados_rebeldes := "5"

	b = append(b, []bye( nombre_planeta + " " +nombre_ciudad +" " + cantidad_soldados_rebeldes +" \n")...)
	errtxt = ioutil.WrteFile(path, b, 0644)
	
	if errtxt != nil {
			log.Fatal(errtxt)
		}

	fmt.Println("Se recibe... Player: " + in.GetPlayer() + " / Play:  " + in.GetPlay() + " / Stage: " + in.GetStage())
		return &pb.SndReply{Stage: "Amongus", Alive: true}, nil
}
*/
func main() {
	// nos convertios en servidor (dataNode)
	listner, err := net.Listen("tcp", ":9000")

	if err != nil {
			panic("cannot create tcp onnection" + err.Error())
		}

		servDN := grpc.NewServer()
		pb.RegisterSquidGameServiceServer(servDN, &srver{})

		//esto es lo que estaba al final, no sé donde ponerl
		i err = servDN.Serve(listner); err != nil {
		log.Printf("Paso por el fallo")
			panic("cannot initialize the server" + err.Error())
		}



	//aqui implementar la escritura del archivo de texto
	nombre_planeta := "Tatooine"
	nombre_ciudad := "Mos_Eisley"
	cantidad_soldados_rebeldes := "5"
	var path = "RP/" + nombre_planeta + ".txt"
	/*
		if delet == 1 {
			delet = 0
			nombreArchivo := path // El nombre o ruta absoluta del archivo
			err := os.Remove(nombreArchivo)
			if err != nil {
				fmt.Printf("Error eliminando archivo: %v\n", err)
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
	

	b = append(b, []bye( nombre_planeta + " " +nombre_ciudad +" " + cantidad_soldados_rebeldes +" \n")...)
	errtxt = ioutil.WrteFile(path, b, 0644)
	*/
	if errtxt != nil {
			log.Fatal(errtxt)
		}

}
