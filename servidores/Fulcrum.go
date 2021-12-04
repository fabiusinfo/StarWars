//serivodes fulcrum son bakanes aaaaa POR LA CONCHETUMAREEEE ahora deberia poderse vamo chupete suazo
package main

import (
	//"context"
	"bufio"
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

type VectorClock struct {
	planet string
	X      int
	Y      int
	Z      int
}

//Global Variables
var VectorClock_list []VectorClock

var ip1, ip2 string

func (s *server) Identify(ctx context.Context, in *pb.SendIp) (*pb.IpRecieve, error) {

	ip1 = in.GetIp1()
	ip2 = in.GetIp2()

	return &pb.IpRecieve{Message: "recibido"}, nil
}

/*func (s *server) SendInformationF(ctx context.Context, in *pb.SendRequest) (*pb.SendReply2, error) {

	return &pb.SendReply2{Message: "Fulcrum recibió tu información con éxito"}, nil
}*/

func (s *server) ConsultPlanet(ctx context.Context, in *pb.ConsultRequest) (*pb.ConsultReply, error) {
	//command := in.GetCommand()
	planet := in.GetPlanet()
	city := in.GetCity()
	soldiers := ""
	var path = "servidores/RP/" + planet + ".txt"
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		if strings.Contains(line, city) {
			splitLine := strings.Split(string(line), " ")
			soldiers = splitLine[2]

		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	aux := 0
	for i := 0; i < len(VectorClock_list); i++ {
		if VectorClock_list[i].planet == planet {
			aux = i
			break

		}
	}

	return &pb.ConsultReply{Rebelds: soldiers, Clock: strconv.Itoa(VectorClock_list[aux].X) + " " + strconv.Itoa(VectorClock_list[aux].Y) + " " + strconv.Itoa(VectorClock_list[aux].Z)}, nil
}

// Crear archivo
func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

// Borrar elemento de array en posicion index
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func crearArchivo(path string, planet string) {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
		VectorClock_list = append(VectorClock_list, VectorClock{planet, 0, 0, 0})
	}
}

func crearArchivo_log(path string) {
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

func (s *server) SendInformationF(ctx context.Context, in *pb.SendRequestF) (*pb.SendReplyF, error) {

	//aqui implementar la escritura del archivo de texto
	command := in.GetCommand()
	planet := in.GetPlanet()
	city := in.GetCity()
	value := in.GetValue()
	var path = "servidores/RP/" + planet + ".txt"

	var path_log = "servidores/RP/log_" + planet + ".txt"

	if command == "DeleteCity" {
		fmt.Println("Comando recibido: " + command + " " + planet + " " + city)
	} else {
		fmt.Println("Comando recibido: " + command + " " + planet + " " + city + " " + value)
	}

	crearArchivo(path, planet)
	crearArchivo_log(path_log)

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

		// añadir al log
		bl, errtxtl := ioutil.ReadFile(path_log)

		if errtxtl != nil {
			log.Fatal(errtxtl)
		}

		bl = append(bl, []byte(command+" "+planet+" "+city+" "+value+" \n")...)
		errtxtl = ioutil.WriteFile(path_log, bl, 0644)

		if errtxtl != nil {
			log.Fatal(errtxtl)
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

		// añadir al log
		bl, errtxtl := ioutil.ReadFile(path_log)

		if errtxtl != nil {
			log.Fatal(errtxtl)
		}

		bl = append(bl, []byte(command+" "+planet+" "+city+" "+value+" \n")...) //value es la nueva ciudd
		errtxtl = ioutil.WriteFile(path_log, bl, 0644)

		if errtxtl != nil {
			log.Fatal(errtxtl)
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

		// añadir al log
		bl, errtxtl := ioutil.ReadFile(path_log)

		if errtxtl != nil {
			log.Fatal(errtxtl)
		}

		bl = append(bl, []byte(command+" "+planet+" "+city+" "+value+" \n")...)
		errtxtl = ioutil.WriteFile(path_log, bl, 0644)

		if errtxtl != nil {
			log.Fatal(errtxtl)
		}
	} else { //DeleteCity
		input, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalln(err)
		}

		lines := strings.Split(string(input), "\n")

		for i, line := range lines {
			if strings.Contains(line, city) {
				lines = RemoveIndex(lines, i)
				break
			}

		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(path, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
		// añadir al log
		bl, errtxtl := ioutil.ReadFile(path_log)

		if errtxtl != nil {
			log.Fatal(errtxtl)
		}

		bl = append(bl, []byte(command+" "+planet+" "+city+" \n")...)
		errtxtl = ioutil.WriteFile(path_log, bl, 0644)

		if errtxtl != nil {
			log.Fatal(errtxtl)
		}
	}
	aux := 0
	for i := 0; i < len(VectorClock_list); i++ {
		if VectorClock_list[i].planet == planet {
			if in.GetFulcrum() == "1" {
				VectorClock_list[i].X += 1
			} else if in.GetFulcrum() == "2" {
				VectorClock_list[i].Y += 1
			} else {
				VectorClock_list[i].Z += 1
			}
			aux = i
			break

		}
	}

	return &pb.SendReplyF{Clock: strconv.Itoa(VectorClock_list[aux].X) + " " + strconv.Itoa(VectorClock_list[aux].Y) + " " + strconv.Itoa(VectorClock_list[aux].Z)}, nil
}



func propagation(){

	for i := 0; i < len(VectorClock_list); i++ {
		readFile, err := os.Open("RP/log_"+ VectorClock_list[i].planet)
		if err != nil {
				log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var lines []string   // aqui se guardan las lineas
		for fileScanner.Scan() {
				lines = append(lines, fileScanner.Text())
		}
		readFile.Close()
		for _, line := range lines {
				og_command := strings.Split(line, " ") //separa el comando en :accion que realiza, planeta, ciudad, y valor (que puede ser nuevo nombre de ciudad o numero de solados)
				command := og_command[0]  
				planet := og_command[1]  
				city := og_command[2]
				if (command == "DeleteCity"){
					value := 0
				}else{
					value := og_command[3]
				}

		}
	
	}
	


}



func main() {
//nos convertios en servidor
	//VectorClock := [3]int{0, 0, 0} //{f1-42, f-43, f3-44}
	//VectorClock  append(VectorClock, 0, 0, 0)

	X := "none"
	go func() {
		listener, err := net.Listen("tcp", ":9000")

		if err != nil {
			panic("cannot connect with servr " + err.Error())
		}

		serv := grpc.NewServer()
		pb.RegisterStarWarsServiceServer(serv, &server{})
		if err = serv.Serve(listener); err != nil {
			panic("cannot initialize the server" +err.Error())

		}

	}()

fmt.Println("<Servidor Fulcrum habilitado>")
	fmt.Scanln(&X)

}
