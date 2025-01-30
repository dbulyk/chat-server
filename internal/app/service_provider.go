package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"chat_server/internal/api/chat"
	"chat_server/internal/closer"
	"chat_server/internal/config"
	"chat_server/internal/config/env"
	"chat_server/internal/repository"
	repo "chat_server/internal/repository/chat"
	"chat_server/internal/service"
	chatService "chat_server/internal/service/chat"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig
	pgConfig   config.PGConfig
	pgPool     *pgxpool.Pool

	chatRepository repository.ChatRepository
	chatService    service.ChatService
	implementation *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) GRPCConfig() config.GRPCConfig {
	if sp.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to load grpc config: %v", err)
		}

		sp.grpcConfig = cfg
	}
	return sp.grpcConfig
}

func (sp *serviceProvider) PGConfig() config.PGConfig {
	if sp.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to load pg config: %v", err)
		}
		sp.pgConfig = cfg
	}
	return sp.pgConfig
}

func (sp *serviceProvider) PGPool(ctx context.Context) *pgxpool.Pool {
	if sp.pgPool == nil {
		conn, err := pgxpool.New(ctx, sp.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = conn.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database: %v", err)
		}

		closer.Add(func() error {
			conn.Close()
			return nil
		})

		sp.pgPool = conn
	}
	return sp.pgPool
}

func (sp *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if sp.chatRepository == nil {
		chatRepo := repo.NewRepository(sp.PGPool(ctx))
		sp.chatRepository = chatRepo
	}
	return sp.chatRepository
}

func (sp *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if sp.chatService == nil {
		serv := chatService.NewChatService(sp.ChatRepository(ctx))
		sp.chatService = serv
	}
	return sp.chatService
}

func (sp *serviceProvider) ChatImplementation(ctx context.Context) *chat.Implementation {
	if sp.implementation == nil {
		impl := chat.NewImplementation(sp.ChatService(ctx))
		sp.implementation = impl
	}
	return sp.implementation
}
