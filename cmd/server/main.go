package main

import (
	"log"

	"github.com/rsh456/go-grpc-services/internal/db"
	"github.com/rsh456/go-grpc-services/internal/rocket"
)

func Run() error {
	//Responsible for initializing and starting
	//gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
	}
	_ = rocket.New(rocketStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
