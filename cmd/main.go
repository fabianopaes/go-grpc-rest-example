package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fabianopaes/go-grpc-rest-example/api"
	"github.com/fabianopaes/go-grpc-rest-example/pkg/repository"
	"github.com/fabianopaes/go-grpc-rest-example/pkg/repository/memory"
	"github.com/fabianopaes/go-grpc-rest-example/pkg/service"
	"github.com/fabianopaes/go-grpc-rest-example/server"
	"google.golang.org/grpc"
)

func main() {
	personRepo := getPersonRepository()
	personGRPC := getPersonServer(personRepo)

	grpc := GetGRPCServerReady(personGRPC)
	grpc.Start()

	rest := GetRestServerReady()
	rest.Start()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan

	rest.Shutdown()
	grpc.Shutdown()

}

//GetGRPCServerReady start the grpc server
func GetGRPCServerReady(pss api.PersonServiceServer) *server.GRPC {
	//FIXME should read the port value from env var
	return server.NewGRPC("5000", pss)
}

//GetRestServerReady gets the http rest server running
func GetRestServerReady() *server.Rest {
	//FIXME should read the port value from env var
	client := getPersonClient()
	return server.NewRest("5001", client)
}

func getPersonRepository() *memory.PersonRepository {
	return memory.NewPerson()
}

func getPersonServer(repo repository.Person) api.PersonServiceServer {
	return service.NewPerson(repo)
}

func getPersonClient() api.PersonServiceClient {
	//FIXME should read the port value from env var
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return api.NewPersonServiceClient(conn)
}
