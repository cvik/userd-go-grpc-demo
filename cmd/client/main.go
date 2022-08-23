package main

import (
	"context"
	"log"

	usersv1 "github.com/cvik/userd/pkg/services/users/v1"
)

func main() {
	log.Println("Creating client to 'lendo.users.v1.Users' service...")
	client, err := usersv1.NewClient("localhost", 1998, "testClient")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Calling 'CreateUser' endpoint..")
	ctx := context.Background()
	createResp, err := client.CreateUser(ctx, &usersv1.CreateUserRequest{
		Name:  "cvi",
		Email: "cvi@heya.tech",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CreateUser created user with ID: %#v\n", createResp.GetId())

	log.Println("Calling 'GetUser' to retrieve user...")
	getResp, err := client.GetUser(ctx, &usersv1.GetUserRequest{Id: createResp.GetId()})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("GetUser found user: %v (%v)\n",
		getResp.GetUser().GetName(), getResp.GetUser().GetEmail())
}
