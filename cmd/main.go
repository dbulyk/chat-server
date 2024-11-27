package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	chatAPI "chat_server/internal/api/chat"
	"chat_server/internal/config"
	"chat_server/internal/config/env"
	"chat_server/internal/repository/chat"
	serv "chat_server/internal/service/chat"
	desc "chat_server/pkg/chat_server_v1"
)

var configPath string

// init записывает параметр конфига
func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

// main запускает сервер на указанном в конфиге порту
func main() {
	flag.Parse()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	chatRepo := chat.NewRepository(pool)
	chatService := serv.NewChatService(chatRepo)
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServerV1Server(s, chatAPI.NewImplementation(chatService))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
