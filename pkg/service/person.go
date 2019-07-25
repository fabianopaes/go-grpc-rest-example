package service

import (
	"context"
	"fmt"

	"github.com/fabianopaes/go-grpc-rest-example/api"
	"github.com/fabianopaes/go-grpc-rest-example/pkg/repository"
)

type Person struct {
	repo repository.Person
}

func NewPerson(repo repository.Person) *Person {
	return &Person{
		repo: repo,
	}
}

func (ps *Person) Get(ctx context.Context, request *api.GetPersonRequest) (*api.Person, error) {
	if request == nil {
		return nil, fmt.Errorf("you must ask by a person. Just let me konw who you wanna get")
	}
	return ps.repo.Get(ctx, request.Id)
}
