package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "github.com/fabiusinfo/StarWars/proto"
	amqp "github.com/rabbitmq/amqp091-go"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStarWarsServiceServer
}

type PlayerStruct struct {
	id    string
	alive bool
	score int
}

var liderPlay int
var actualStage string
var actualRound int32
var started bool
var flaggy bool
var list_of_players []PlayerStruct

//listas stage 2
var group1 []PlayerStruct
var group2 []PlayerStruct
var groupaux []PlayerStruct

//listas stage 3
var group3 []PlayerStruct

//var players [16]string
var totalPlayers int

// Error para el Rabbit
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (s *server) JoinGame(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	//players[in.GetPlayer()] = "alive"
	totalPlayers += 1
	list_of_players = append(list_of_players, PlayerStruct{in.GetPlayer(), true, 0})
	return &pb.JoinReply{Message: "inscrito con exito"}, nil
}

func (s *server) Started(ctx context.Context, in *pb.StartRequest) (*pb.StartReply, error) {
	return &pb.StartReply{Started: started}, nil
}

func (s *server) AskRound(ctx context.Context, in *pb.AskRequest) (*pb.AskReply, error) {
	return &pb.AskReply{Round: int32(actualRound)}, nil
}

func (s *server) DeadOrAlive(ctx context.Context, in *pb.DeadRequest) (*pb.DeadReply, error) {
	alive := true
	if in.GetStage() == "1rv" {
		for i := 0; i < 16; i++ {
			if list_of_players[i].id == in.GetPlayer() {
				alive = list_of_players[i].alive
			}
		}
	} else if in.GetStage() == "2tc" {
		for i := 0; i < len(group1); i++ {
			if group1[i].id == in.GetPlayer() {
				alive = group1[i].alive
			}
		}
		for i := 0; i < len(group2); i++ {
			if group2[i].id == in.GetPlayer() {
				alive = group2[i].alive
			}
		}
	} else {
		log.Printf("para el nivel 3 no es necesario")

	}

	return &pb.DeadReply{Dead: alive}, nil
}

func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	alive := true
	flaggy = true

	if actualRound != 0 {
		if in.GetRound() == actualRound {

			//envío al nameNode
			conn, err := grpc.Dial("10.6.43.42:8080", grpc.WithInsecure())

			if err != nil {
				panic("cannot connect with server " + err.Error())
			}

			serviceLider := pb.NewStarWarsServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			_, err = serviceLider.SendPlays(ctx, &pb.SendRequest{Player: in.GetPlayer(), Play: in.GetPlay(), Stage: in.GetStage(), Round: in.GetRound()})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
		
			//Envío al Pozo

			if started == true {
				pPlay, errpPlay := strconv.Atoi(in.GetPlay())
				if errpPlay != nil {
					log.Fatalf("could not greet: %v", errpPlay)
				}
				if actualStage == "1rv" {
					for i := 0; i < 16; i++ {
						if list_of_players[i].id == in.GetPlayer() {

							list_of_players[i].score += pPlay
						}
					}
				} else if actualStage == "2tc" {
					for i := 0; i < len(group1); i++ {
						if group1[i].id == in.GetPlayer() {

							group1[i].score += pPlay
						}
					}
					for i := 0; i < len(group2); i++ {
						if group2[i].id == in.GetPlayer() {

							group2[i].score += pPlay
						}
					}
				} else {
					for i := 0; i < len(group3); i++ {
						if group3[i].id == in.GetPlayer() {

							group3[i].score += pPlay
						}
					}
				}

				if actualStage == "1rv" {

					if pPlay > liderPlay {
						alive = false
						for i := 0; i < 16; i++ {
							if list_of_players[i].id == in.GetPlayer() {
								list_of_players[i].alive = false
							}
						}
						conn, err := amqp.Dial("amqp://admin:test@10.6.43.41:5672/")
						failOnError(err, "Failed to connect to RabbitMQ")
						defer conn.Close()

						ch, err := conn.Channel()
						failOnError(err, "Failed to open a channel")
						defer ch.Close()

						q, err := ch.QueueDeclare(
							"hello", // name
							false,   // durable
							false,   // delete when unused
							false,   // exclusive
							false,   // no-wait
							nil,     // arguments
						)
						failOnError(err, "Failed to declare a queue")

						i := in.GetPlayer()
						s := in.GetStage()

						body := "Jugador_" + i + " Ronda_" + s

						err = ch.Publish(
							"",     // exchange
							q.Name, // routing key
							false,  // mandatory
							false,  // immediate
							amqp.Publishing{
								ContentType: "text/plain",
								Body:        []byte(body),
							})
						failOnError(err, "Failed to publish a message")
						log.Printf("Ha muerto: %s ", body)
					}
				}
			} else {
				log.Printf("aún no comienza el nivel")

				return &pb.SendReply{Stage: actualStage, Alive: alive, Round: in.GetRound()}, nil
			}
			return &pb.SendReply{Stage: actualStage, Alive: alive, Round: in.GetRound() + 1}, nil
		} else {
			log.Printf("ya realizaste la jugada en esta ronda ")
			return &pb.SendReply{Stage: actualStage, Alive: alive, Round: in.GetRound()}, nil
		}

	} else {
		log.Printf("el lider todavía no comienza la ronda")
		return &pb.SendReply{Stage: actualStage, Alive: alive, Round: in.GetRound()}, nil
	}

}



func main() {
	//códigos Etapas
	//1rv
	//2tc
	//3tn

	go func() {
		// nos convertimos en servidor (LIDER)
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


}