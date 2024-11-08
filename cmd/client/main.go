package main

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	desc "chat_server/pkg/chat_server_v1"
)

const (
	address = "localhost:50051"
)

// main создает соединение с сервером
func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			slog.Error("conn isn't closed", "error", err.Error())
		}
	}(conn)

	c := desc.NewChatServerV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := c.CreateChat(ctx, &desc.CreateChatRequest{
		UserIds: []int64{1, 2},
		Title:   gofakeit.Name(),
	})
	if err != nil {
		log.Panicf("failed to create chat by id: %v", err)
	}

	log.Printf(color.RedString("Chat id:\n"), color.GreenString("%+v", r.GetId()))
}
