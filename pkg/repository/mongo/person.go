package mongo

import (
	"context"
	"fmt"

	"github.com/fabianopaes/go-grpc-rest-example/api"
	"github.com/globalsign/mgo"
)

type PersonRepository struct {
	session *mgo.Session
}

func NewPersonRepository(ses *mgo.Session) *PersonRepository {
	return &PersonRepository{
		session: ses,
	}
}

func (pr *PersonRepository) Get(ctx context.Context, ID string) (*api.Person, error) {
	return nil, fmt.Errorf("not implemented yet")
}
