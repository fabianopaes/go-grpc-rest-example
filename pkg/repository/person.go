package repository

import (
	"context"

	"github.com/fabianopaes/go-grpc-rest-example/api"
)

type Person interface {
	Get(ctx context.Context, ID string) (*api.Person, error)
}
