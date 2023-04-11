package main

import (
	"context"
	"fmt"
	"github.com/grpc_server/server/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
	}

	client := user.NewInventoryClient(conn)
	resp, err := client.GetUserList(context.Background(), &user.GetUserListRequest{
		Name: "h",
	})
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Printf("%+v", resp)
}
