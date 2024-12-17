package app

import (
	"context"
	"log"

	"chat_server/internal/api/chat"
	"chat_server/internal/client/db"
	"chat_server/internal/client/db/pg"
	"chat_server/internal/closer"
	"chat_server/internal/config"
	"chat_server/internal/repository"
	chatRepository "chat_server/internal/repository/chat"
	"chat_server/internal/service"
	chatService "chat_server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient db.Client
	//txManager      db.TxManager
	chatRepository repository.ChatServerRepository

	chatService service.ChatService

	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

//
//func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
//	if s.txManager == nil {
//		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
//	}
//
//	return s.txManager
//}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatServerRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewChatService(s.ChatRepository(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) NoteImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
