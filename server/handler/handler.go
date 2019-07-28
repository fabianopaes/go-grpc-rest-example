package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fabianopaes/go-grpc-rest-example/api"
	"github.com/julienschmidt/httprouter"
)

//New creates the handler to serve http requests
func New(client api.PersonServiceClient) http.Handler {
	log.Println("creating the http handlers. It will be only the proxy to grpc")

	router := httprouter.New()

	router.GET("/persons", getPersons(client))
	router.GET("/persons/:personId", getPerson(client))

	log.Println("http handlers(proxy) has been created")
	return router
}

//GetPerson returns the person
func getPerson(person api.PersonServiceClient) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// fmt.Fprint(w, "Get a single person!\n")
		request := &api.GetPersonRequest{
			Id: p.ByName("personId"),
		}
		result, err := person.Get(context.Background(), request)
		if err != nil {
			log.Println(err)
			http.Error(w, "the person does not exists.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

//GetPersons returns the list of persons
func getPersons(client api.PersonServiceClient) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Get a single person!\n")
	}
}
