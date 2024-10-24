package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "chat_server/pkg/chat_server_v1"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatServerV1Server
}

func (s *server) Create(_ context.Context, in *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Received CreateRequest: %s", in.GetName())
	return &desc.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (s *server) Delete(_ context.Context, in *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Received DeleteRequest: %d", in.GetId())
	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, in *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Received SendMessageRequest: %v", in)
	return &emptypb.Empty{}, nil
}

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
