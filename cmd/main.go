package main

import (
	"github.com/fabianopaes/go-grpc-rest-example/api"
	"github.com/fabianopaes/go-grpc-rest-example/pkg/repository"
	"github.com/fabianopaes/go-grpc-rest-example/pkg/repository/memory"
	"github.com/fabianopaes/go-grpc-rest-example/pkg/service"
	"github.com/fabianopaes/go-grpc-rest-example/server"
)

func main() {
	personRepo := getPersonRepository()
	personGRPC := getPersonServer(personRepo)
	grpc := GetGRPCServerReady(personGRPC)
	grpc.Start()

	// rest := GetRestServerReady()
	// rest.Start()
}

//GetGRPCServerReady start the grpc server
func GetGRPCServerReady(pss api.PersonServiceServer) server.Server {
	return server.NewGRPC("5000", pss)
}

//GetRestServerReady gets the http rest server running
func GetRestServerReady() server.Server {
	return &server.Rest{}
}

func getPersonRepository() *memory.PersonRepository {
	return &memory.PersonRepository{}
}

func getPersonServer(repo repository.Person) api.PersonServiceServer {
	return service.NewPerson(repo)
}
