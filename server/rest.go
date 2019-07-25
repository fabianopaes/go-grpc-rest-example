package server

import "log"

type Rest struct {
}

func NewRest() *Rest {
	return &Rest{}
}

func (r *Rest) Start() {
	log.Println("just to write down")
}
