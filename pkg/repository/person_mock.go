package repository

import (
	"context"

	"github.com/fabianopaes/go-grpc-rest-example/api"
)

type PersonMock struct {
	GetFn             func(ctx context.Context, ID string) (*api.Person, error)
	GetFnInvokedCount int
}

func (pm *PersonMock) Get(ctx context.Context, ID string) (*api.Person, error) {
	pm.GetFnInvokedCount++
	return pm.GetFn(ctx, ID)
}
