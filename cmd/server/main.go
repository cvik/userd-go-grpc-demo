package main

import (
	"log"

	usersv1 "github.com/cvik/userd-go-grpc-demo/pkg/services/users/v1"
	"github.com/cvik/userd-go-grpc-demo/pkg/store"
)

func main() {
	memStore := store.NewMemStore()

	log.Println("Running gRPC server on port 1998...")
	err := usersv1.NewServer(memStore).Run(1998)
	if err != nil {
		log.Fatal(err)
	}
}
