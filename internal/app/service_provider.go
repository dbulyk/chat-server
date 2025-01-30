package app

import (
	"context"
	"log"

	"chat_server/internal/api/chat"
	"chat_server/internal/client/db"
	"chat_server/internal/client/db/pg"
	"chat_server/internal/client/db/trancsation"
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
	dbc        db.Client
	txManager  db.TxManager

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

func (sp *serviceProvider) DBClient(ctx context.Context) db.Client {
	if sp.dbc == nil {
		conn, err := pg.New(ctx, sp.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = conn.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database: %v", err)
		}

		closer.Add(func() error {
			err = conn.Close()
			if err != nil {
				return err
			}
			return nil
		})

		sp.dbc = conn
	}
	return sp.dbc
}

func (sp *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if sp.txManager == nil {
		sp.txManager = trancsation.NewTransactionManager(sp.DBClient(ctx).DB())
	}

	return sp.txManager
}

func (sp *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if sp.chatRepository == nil {
		chatRepo := repo.NewRepository(sp.DBClient(ctx))
		sp.chatRepository = chatRepo
	}
	return sp.chatRepository
}

func (sp *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if sp.chatService == nil {
		serv := chatService.NewChatService(sp.ChatRepository(ctx), sp.TxManager(ctx))
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
