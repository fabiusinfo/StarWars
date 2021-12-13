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
	"time"

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
	X      int32
	Y      int32
	Z      int32
}

//Global Variables
var VectorClock_list []VectorClock
var GlobalCont int32
var ip, ip1, ip2 string

func (s *server) Identify(ctx context.Context, in *pb.SendIp) (*pb.IpRecieve, error) {
	ip = in.GetIp()
	ip1 = in.GetIp1()
	ip2 = in.GetIp2()

	return &pb.IpRecieve{Message: "recibido"}, nil
}

func (s *server) FulcrumComunication(ctx context.Context, in *pb.CommandsRequest) (*pb.CommandsReply, error) {
	text_t := in.GetCommands()
	cont := in.GetCont()
	//clock := in.GetClock() //esto es x y z

	//acá leer el texto y escribir en el log de registros y archivo de texto de los planetas
	text := strings.Split(text_t, "|")
	aux := 0

	/*
		var lines []string // aqui se guardan las lineas

		for text.Scan() {
			lines = append(lines, text.Text())
		}
	*/

	for _, line := range text {
		og_command := strings.Split(line, " ") //separa el comando en :accion que realiza, planeta, ciudad, y valor (que puede ser nuevo nombre de ciudad o numero de solados)
		command := og_command[0]
		planet := og_command[1]
		city := og_command[2]

		var path = "servidores/RP/" + planet + ".txt"
		var path_log = "servidores/RP/log_" + planet + ".txt"
		crearArchivo(path, planet)
		crearArchivo_log(path_log)

		for i := 0; i < len(VectorClock_list); i++ {
			aux = i
			if VectorClock_list[i].planet == planet {
				if VectorClock_list[i].X < in.GetX() {
					VectorClock_list[i].X = in.GetX()
				}
				if VectorClock_list[i].Y < in.GetY() {
					VectorClock_list[i].Y = in.GetY()
				}
				if VectorClock_list[i].Z < in.GetZ() {
					VectorClock_list[i].Z = in.GetZ()
				}

			}
		}

		if command == "DeleteCity" {
			fmt.Println("Comando recibido: " + command + " " + planet + " " + city)
		} else {
			value := og_command[3]
			fmt.Println("Comando recibido: " + command + " " + planet + " " + city + " " + value)
		}

		if command == "AddCity" {
			// añadir al texto
			b, errtxt := ioutil.ReadFile(path)

			if errtxt != nil {
				log.Fatal(errtxt)
			}
			value := og_command[3]
			b = append(b, []byte(planet+" "+city+" "+value+" \n")...)
			errtxt = ioutil.WriteFile(path, b, 0644)

			if errtxt != nil {
				log.Fatal(errtxt)
			}

			// añadir al log
			if cont != 3 {
				bl, errtxtl := ioutil.ReadFile(path_log)

				if errtxtl != nil {
					log.Fatal(errtxtl)
				}

				bl = append(bl, []byte(command+" "+planet+" "+city+" "+value+" \n")...)
				errtxtl = ioutil.WriteFile(path_log, bl, 0644)

				if errtxtl != nil {
					log.Fatal(errtxtl)
				}

			}

		} else if command == "UpdateName" {

			input, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatalln(err)
			}

			lines := strings.Split(string(input), "\n")
			value := og_command[3]
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

			if cont != 3 {
				bl, errtxtl := ioutil.ReadFile(path_log)

				if errtxtl != nil {
					log.Fatal(errtxtl)
				}

				bl = append(bl, []byte(command+" "+planet+" "+city+" "+value+" \n")...)
				errtxtl = ioutil.WriteFile(path_log, bl, 0644)

				if errtxtl != nil {
					log.Fatal(errtxtl)
				}

			}

		} else if command == "UpdateNumber" {
			input, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatalln(err)
			}

			lines := strings.Split(string(input), "\n")
			value := og_command[3]
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

			if cont != 3 {
				bl, errtxtl := ioutil.ReadFile(path_log)

				if errtxtl != nil {
					log.Fatal(errtxtl)
				}

				bl = append(bl, []byte(command+" "+planet+" "+city+" "+value+" \n")...)
				errtxtl = ioutil.WriteFile(path_log, bl, 0644)

				if errtxtl != nil {
					log.Fatal(errtxtl)
				}

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
			if cont != 3 {
				bl, errtxtl := ioutil.ReadFile(path_log)

				if errtxtl != nil {
					log.Fatal(errtxtl)
				}

				bl = append(bl, []byte(command+" "+planet+" "+city+"\n")...)
				errtxtl = ioutil.WriteFile(path_log, bl, 0644)

				if errtxtl != nil {
					log.Fatal(errtxtl)
				}

			}
		}
	}
	//crear mensaje que se enviará al siguiente fulcrum
	if cont != 3 {
		for i := 0; i < len(VectorClock_list); i++ {
			readFile, err := os.Open("servidores/RP/log_" + VectorClock_list[i].planet + ".txt")
			if err != nil {
				log.Fatal(err)
			}
			fileScanner := bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)
			var lines []string
			for fileScanner.Scan() {
				lines = append(lines, fileScanner.Text())
			}

			readFile.Close()
			commands_strings := ""
			for _, line := range lines {
				commands_strings += line + "|"
			}
			//todos los comandos se los mando al fulcrum que corresponde
			if ip == "10.6.43.42" {

				conn, err := grpc.Dial("10.6.43.43:9000", grpc.WithInsecure())

				if err != nil {
					panic("cannot connect with server " + err.Error())
				}
				servicePropagation := pb.NewStarWarsServiceClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				_, errr := servicePropagation.FulcrumComunication(ctx, &pb.CommandsRequest{Commands: commands_strings, Cont: in.GetCont() + 1, X: VectorClock_list[aux].X, Y: VectorClock_list[aux].Y, Z: VectorClock_list[aux].Z})
				if err != nil {
					log.Fatalf("could not greet: %v", errr)
				}

			} else if ip == "10.6.43.43" {
				conn, err := grpc.Dial("10.6.43.44:9000", grpc.WithInsecure())

				if err != nil {
					panic("cannot connect with server " + err.Error())
				}
				servicePropagation := pb.NewStarWarsServiceClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				_, errr := servicePropagation.FulcrumComunication(ctx, &pb.CommandsRequest{Commands: commands_strings, Cont: in.GetCont() + 1, X: VectorClock_list[aux].X, Y: VectorClock_list[aux].Y, Z: VectorClock_list[aux].Z})
				if err != nil {
					log.Fatalf("could not greet: %v", errr)
				}

			} else {
				var ipe string
				for i = 0; i < 2; i++ {
					if i == 0 {
						ipe = "10.6.43.42"
					} else {
						ipe = "10.6.43.43"
					}
					conn, err := grpc.Dial(ipe+":9000", grpc.WithInsecure())

					if err != nil {
						panic("cannot connect with server " + err.Error())
					}
					servicePropagation := pb.NewStarWarsServiceClient(conn)

					ctx, cancel := context.WithTimeout(context.Background(), time.Second)
					defer cancel()

					_, errr := servicePropagation.FulcrumComunication(ctx, &pb.CommandsRequest{Commands: commands_strings, Cont: in.GetCont() + 1, X: VectorClock_list[aux].X, Y: VectorClock_list[aux].Y, Z: VectorClock_list[aux].Z})
					if err != nil {
						log.Fatalf("could not greet: %v", errr)
					}
				}
			}

			// -ESTO FALTA acá vaciar archivo de texto y log de registro solo si es fulcrum 1 y 2
			if cont != 3 {
				if ip == "10.6.43.42" {
					for i := 0; i < len(VectorClock_list); i++ {
						//se borra el archivo log del planeta y archivo planeta del fulcrum 1
						file_log := os.Remove("servidores/RP/log_" + VectorClock_list[i].planet + ".txt")
						if file_log != nil {
							log.Fatal(file_log)
						}

						file_planet := os.Remove(VectorClock_list[i].planet + ".txt")
						if file_planet != nil {
							log.Fatal(file_planet)
						}
					}
				} else if ip == "10.6.43.43" {
					for i := 0; i < len(VectorClock_list); i++ {
						//se borra el archivo log del planeta y archivo planeta del fulcrum 2
						file_log := os.Remove("servidores/RP/log_" + VectorClock_list[i].planet + ".txt")
						if file_log != nil {
							log.Fatal(file_log)
						}

						file_planet := os.Remove(VectorClock_list[i].planet + ".txt")
						if file_planet != nil {
							log.Fatal(file_planet)
						}
					}
				}
			}
		}
	}

	return &pb.CommandsReply{Message: "Fulcrum recibió tu información con éxito"}, nil
}

func (s *server) ConsultPlanet(ctx context.Context, in *pb.ConsultRequest) (*pb.ConsultReply, error) {
	//command := in.GetCommand()
	planet := in.GetPlanet()
	city := in.GetCity()
	soldiers := "none"
	VectorClock := "none"
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
			VectorClock = strconv.Itoa(int(VectorClock_list[aux].X)) + " " + strconv.Itoa(int(VectorClock_list[aux].Y)) + " " + strconv.Itoa(int(VectorClock_list[aux].Z))
			break

		}
	}

	return &pb.ConsultReply{Rebelds: soldiers, Clock: VectorClock}, nil
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
			if ip1 == "10.6.43.42" {
				VectorClock_list[i].X += 1
			} else if ip1 == "10.6.43.43" {
				VectorClock_list[i].Y += 1
			} else {
				VectorClock_list[i].Z += 1
			}
			aux = i
			break

		}
	}

	return &pb.SendReplyF{Clock: strconv.Itoa(int(VectorClock_list[aux].X)) + " " + strconv.Itoa(int(VectorClock_list[aux].Y)) + " " + strconv.Itoa(int(VectorClock_list[aux].Z))}, nil
}

func propagation() {
	var ipe string
	var value string
	for i := 0; i < len(VectorClock_list); i++ {
		readFile, err := os.Open("servidores/RP/log_" + VectorClock_list[i].planet + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var lines []string // aqui se guardan las lineas
		for fileScanner.Scan() {
			lines = append(lines, fileScanner.Text())
		}
		readFile.Close()
		for _, line := range lines {
			og_command := strings.Split(line, " ") //separa el comando en :accion que realiza, planeta, ciudad, y valor (que puede ser nuevo nombre de ciudad o numero de solados)
			command := og_command[0]
			planet := og_command[1]
			city := og_command[2]

			if command == "DeleteCity" {
				value = "0"
			} else {
				value = og_command[3]
			}
			for i := 0; i < 2; i++ {
				if i == 0 {
					ipe = ip1
				} else {
					ipe = ip2
				}

				conn, err := grpc.Dial(ipe+":9000", grpc.WithInsecure())

				if err != nil {
					panic("cannot connect with server " + err.Error())
				}
				servicePropagation := pb.NewStarWarsServiceClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				_, errr := servicePropagation.SendInformationF(ctx, &pb.SendRequestF{Command: command, Planet: planet, City: city, Value: value})
				if err != nil {
					log.Fatalf("could not greet: %v", errr)
				}
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
			panic("cannot initialize the server" + err.Error())

		}

	}()

	go func() {

		for true {
			timer := time.NewTimer(2 * time.Minute)
			<-timer.C
			fmt.Println("pasaron 2 min")
			//fmt.Scanln(&X)

			//esto cada 2 min
			if ip == "10.6.43.42" {
				//leer log de registro, vaciar log y planet.txt y enviar en commands
				for i := 0; i < len(VectorClock_list); i++ {
					readFile, err := os.Open("servidores/RP/log_" + VectorClock_list[i].planet + ".txt")
					if err != nil {
						log.Fatal(err)
					}
					fileScanner := bufio.NewScanner(readFile)
					fileScanner.Split(bufio.ScanLines)
					var lines []string
					for fileScanner.Scan() {
						lines = append(lines, fileScanner.Text())
					}

					readFile.Close()
					commands_strings := ""
					for _, line := range lines {
						commands_strings += line + "|"
					}
					//todos los comandos se los mando al fulcrum 2

					conn, err := grpc.Dial("10.6.43.43:9000", grpc.WithInsecure())

					if err != nil {
						panic("cannot connect with server " + err.Error())
					}
					servicePropagation := pb.NewStarWarsServiceClient(conn)

					ctx, cancel := context.WithTimeout(context.Background(), time.Second)
					defer cancel()

					_, errr := servicePropagation.FulcrumComunication(ctx, &pb.CommandsRequest{Commands: commands_strings, Cont: 1, X: VectorClock_list[i].X, Y: VectorClock_list[i].Y, Z: VectorClock_list[i].Z})
					if err != nil {
						log.Fatalf("could not greet: %v", errr)
					}
					readFile.Close()
					//se borra el archivo log del planeta
					file_log := os.Remove("servidores/RP/log_" + VectorClock_list[i].planet + ".txt")
					if file_log != nil {
						log.Fatal(file_log)
					}
					//se borra el archivo del planeta
					file_planet := os.Remove(VectorClock_list[i].planet + ".txt")
					if file_planet != nil {
						log.Fatal(file_planet)
					}
				}

			}
			fmt.Println("paso por acá")

		}

	}()
	fmt.Println("<Servidor Fulcrum habilitado>")
	fmt.Scanln(&X)

}
