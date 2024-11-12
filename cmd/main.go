package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"chat_server/internal/config"
	"chat_server/internal/config/env"
	desc "chat_server/pkg/chat_server_v1"
)

var configPath string

type server struct {
	desc.UnimplementedChatServerV1Server
	db *pgx.Conn
}

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

	conn, err := pgx.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer func(pool *pgx.Conn, ctx context.Context) {
		err = pool.Close(ctx)
		if err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}(conn, ctx)

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServerV1Server(s, &server{db: conn})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// CreateChat Создаёт новый чат с указанными пользователями и названием
func (s *server) CreateChat(ctx context.Context, in *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	builder := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("title").
		Values(in.GetTitle()).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var chatID int64
	err = s.db.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		return nil, err
	}

	builder = sq.Insert("users_chats").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_tag")

	for _, v := range in.GetUsersTags() {
		builder = builder.Values(chatID, v)
	}

	query, args, err = builder.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &desc.CreateChatResponse{ChatId: chatID}, nil
}

// AddUserToChat добавляет пользователей в уже созданный чат
func (s *server) AddUserToChat(ctx context.Context, in *desc.AddUsersToChatRequest) (*emptypb.Empty, error) {
	builder := sq.Insert("users_chats").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_tag")
	for _, v := range in.GetUsersTag() {
		builder = builder.Values(in.GetChatId(), v)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteChat удаляет чат
func (s *server) DeleteChat(ctx context.Context, in *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	builder := sq.Delete("chat").
		Where(sq.Eq{"id": in.GetChatId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// SendMessage отправляет сообщение в чат
func (s *server) SendMessage(ctx context.Context, in *desc.SendMessageRequest) (*emptypb.Empty, error) {
	builder := sq.Insert("messages").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_tag", "message").
		Values(in.GetChatId(), in.GetUserTag(), in.GetText())

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
