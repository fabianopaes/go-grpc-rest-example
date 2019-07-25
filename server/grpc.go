package server

import (
	"fmt"
	"log"
	"net"

	"github.com/fabianopaes/go-grpc-rest-example/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server interface {
	Start()
}

type GRPC struct {
	port    string
	pserver api.PersonServiceServer
}

//NewServer will start the new gRPC serer
func NewGRPC(port string, pserver api.PersonServiceServer) *GRPC {
	return &GRPC{
		port:    port,
		pserver: pserver,
	}
}

//Start starts the gRPC server
func (g *GRPC) Start() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", g.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterPersonServiceServer(s, g.pserver)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
