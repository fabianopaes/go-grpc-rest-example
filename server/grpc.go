package server

import (
	"fmt"
	"log"
	"net"

	"github.com/fabianopaes/go-grpc-rest-example/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Server defines how the application server must looks like
type Server interface {
	Start()
	Shutdown()
}

//GRPC will serve the rpc requests
type GRPC struct {
	port    string
	pserver api.PersonServiceServer
	server  *grpc.Server
}

//NewGRPC will start the new gRPC serer
func NewGRPC(port string, pserver api.PersonServiceServer) *GRPC {
	return &GRPC{
		port:    port,
		pserver: pserver,
		server:  grpc.NewServer(),
	}
}

//Start starts the gRPC server
func (g *GRPC) Start() {
	go func() {
		log.Println("the grpc server is starting up")
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", g.port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		api.RegisterPersonServiceServer(g.server, g.pserver)
		// Register reflection service on gRPC server.
		reflection.Register(g.server)
		log.Println("the grpc server is up")
		if err := g.server.Serve(lis); err != nil {
			log.Fatalf("failed to serve the grpc: %v", err)
		}
	}()
}

//Shutdown makes the graceful shutdown for grpc server
func (g *GRPC) Shutdown() {
	g.server.GracefulStop()
}
