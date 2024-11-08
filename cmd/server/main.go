package main

import (
	desc "chat_server/pkg/chat_server_v1"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/brianvoe/gofakeit"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatServerV1Server
}

// main запускает сервер на указанном порту
func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServerV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// CreateChat Создаёт новый чат с указанными пользователями и названием
func (s *server) CreateChat(_ context.Context, in *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	log.Printf("Received CreateRequest: %s", in.GetTitle())
	return &desc.CreateChatResponse{Id: gofakeit.Int64()}, nil
}

// AddUserToChat добавляет пользователей в уже созданный чат
func (s *server) AddUserToChat(_ context.Context, in *desc.AddUserToChatRequest) (*emptypb.Empty, error) {
	log.Printf("Received AddUserToChatRequest: %v", in.GetUserIds())
	return &emptypb.Empty{}, nil
}

// DeleteChat удаляет чат
func (s *server) DeleteChat(_ context.Context, in *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Printf("Received DeleteRequest: %d", in.GetId())
	return &emptypb.Empty{}, nil
}

// SendMessage отправляет сообщение в чат
func (s *server) SendMessage(_ context.Context, in *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Received SendMessageRequest: %v", in)
	return &emptypb.Empty{}, nil
}
