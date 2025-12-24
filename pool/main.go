package main

import (
	"io"
	"log"
	"pool/pool"
	"sync/atomic"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

// Simulates a resource to share
type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: connection", dbConn.ID)
	return nil
}

// Gives each connection a unique id
var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: new connection", id)

	return &dbConnection{id}, nil
}

func main() {

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	log.Println(p)

}
