package main

import (
	"github.com/fabianopaes/go-grpc-rest-example/api"
	"github.com/fabianopaes/go-grpc-rest-example/pkg/repository"
	"github.com/fabianopaes/go-grpc-rest-example/pkg/repository/memory"
	"github.com/fabianopaes/go-grpc-rest-example/pkg/service"
	"github.com/fabianopaes/go-grpc-rest-example/server"
)

func main() {
	grpc := GetGRPCServerReady()
	grpc.Start()

	rest := GetRestServerReady()
	rest.Start()
}

func GetGRPCServerReady() server.Server {
	return &server.GRPC{}
}

func GetRestServerReady() server.Server {
	return &server.Rest{}
}

func getPersonRepository() *memory.PersonRepository {
	return &memory.PersonRepository{}
}

func getPersonServer(repo *repository.Person) api.PersonServiceServer {
	return service.NewPerson(repo)
}
