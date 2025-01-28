package chat

import (
	"chat_server/internal/repository/chat/model"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	chatTableName      = "chats"
	chatUsersTableName = "chat_users"
	messagesTableName  = "messages"

	idColumn      = "id"
	titleColumn   = "title"
	chatIdColumn  = "chat_id"
	userTagColumn = "user_tag"
	messageColumn = "message"
)

type repo struct {
	db *pgxpool.Pool
}

//func NewRepository(db *pgxpool.Pool) *repository.ChatRepository {
//	return &repo{db: db}
//}

func (r *repo) CreateChat(ctx context.Context, in *model.CreateChatRequest) (*model.CreateChatResponse, error) {
	builder := sq.Insert(chatTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn).
		Values(in.Title).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var chatID int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		return nil, err
	}

	builder = sq.Insert(chatUsersTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIdColumn, userTagColumn)

	for _, v := range in.UserTags {
		builder = builder.Values(chatID, v)
	}

	query, args, err = builder.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &model.CreateChatResponse{ChatId: chatID}, nil
}

func (r *repo) AddUserToChat(ctx context.Context, in *model.AddUserToChatRequest) (*emptypb.Empty, error) {
	builder := sq.Insert(chatUsersTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIdColumn, userTagColumn)
	for _, v := range in.UserTags {
		builder = builder.Values(in.ChatId, v)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (r *repo) DeleteChat(ctx context.Context, in *model.DeleteChatRequest) (*emptypb.Empty, error) {
	builder := sq.Delete(chatUsersTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIdColumn: in.ChatId})
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	builder = sq.Delete(chatTableName).
		Where(sq.Eq{"id": in.ChatId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err = builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (r *repo) SendMessageToChat(ctx context.Context, in *model.SendMessageToChatRequest) (*emptypb.Empty, error) {
	builder := sq.Insert(messagesTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIdColumn, userTagColumn, messageColumn).
		Values(in.ChatId, in.UserTag, in.Message)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
