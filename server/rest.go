package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fabianopaes/go-grpc-rest-example/api"
	"github.com/fabianopaes/go-grpc-rest-example/server/handler"
)

//Rest is the http server
type Rest struct {
	port   string
	server *http.Server
}

//NewRest is the constructor of type Server
func NewRest(port string, client api.PersonServiceClient) *Rest {
	log.Println("creating the http server")
	return &Rest{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%s", port),
			Handler:      handler.New(client),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 55 * time.Second,
		},
	}
}

//Start calls http.server.ListenAndServe
func (r *Rest) Start() {
	log.Println("http server is starting up")
	go func() {
		if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("impossible to start the http rest server")
		}
	}()
}

//Shutdown calls http.server.Shutdown
func (r *Rest) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := r.server.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		return
	}
}
