package main

import (
	"context"
	"fmt"
	"gRPC_Users/pkg"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pkg.NewUserApiClient(conn)

	resp, err := client.CreateUser(context.Background(), &pkg.CreateUserRequest{Name: "Kristina", Email: "fix@mail.ru"})
	if err != nil {
		return
	}

	fmt.Println(resp)

}
