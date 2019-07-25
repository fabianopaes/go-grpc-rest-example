package memory

import (
	"context"
	"fmt"

	"github.com/fabianopaes/go-grpc-rest-example/api"
)

type PersonRepository struct {
	byID   map[string]*api.Person
	byName map[string]*api.Person
}

func NewPerson() *PersonRepository {
	id := map[string]*api.Person{
		"1": {
			Id:   "1",
			Name: "Fabiano 01",
		},
		"2": {
			Id:   "2",
			Name: "Fabiano 02",
		},
		"3": {
			Id:   "3",
			Name: "Fabiano 03",
		},
		"4": {
			Id:   "4",
			Name: "Fabiano 04",
		},
	}

	name := map[string]*api.Person{
		"Fabiano 01": {
			Id:   "1",
			Name: "Fabiano 01",
		},
		"Fabiano 02": {
			Id:   "2",
			Name: "Fabiano 02",
		},
		"Fabiano 03": {
			Id:   "3",
			Name: "Fabiano 03",
		},
		"Fabiano 04": {
			Id:   "4",
			Name: "Fabiano 04",
		},
	}

	return &PersonRepository{
		byID:   id,
		byName: name,
	}
}

func (pr *PersonRepository) Get(ctx context.Context, ID string) (*api.Person, error) {
	per, ok := pr.byID[ID]
	if ok != true {
		return nil, fmt.Errorf("person not found")
	}
	return per, nil
}
