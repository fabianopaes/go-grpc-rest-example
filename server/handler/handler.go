package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fabianopaes/go-grpc-rest-example/api"
	"github.com/julienschmidt/httprouter"
)

//New creates the handler to serve http requests
func New(client *api.PersonServiceClient) http.Handler {
	log.Println("creating the http handlers. It will be only the proxy to grpc")

	router := httprouter.New()

	router.GET("/persons", getPerson(client))
	router.GET("/persons/:personId", getPersons(client))

	log.Println("http handlers(proxy) has been created")
	return router
}

//GetPerson returns the person
func getPerson(client *api.PersonServiceClient) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Get a single person!\n")

	}
}

//GetPersons returns the list of persons
func getPersons(client *api.PersonServiceClient) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Get a single person!\n")

	}
}
