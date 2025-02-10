package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/dbulyk/platform_common/pkg/db"

	"chat_server/internal/model"
	"chat_server/internal/repository"
)

const (
	chatTableName      = "chats"
	chatUsersTableName = "users_chats"
	messagesTableName  = "messages"

	idColumn      = "id"
	titleColumn   = "title"
	chatIDColumn  = "chat_id"
	userTagColumn = "user_tag"
	messageColumn = "message"
)

type repo struct {
	db db.Client
}

// NewRepository возвращает объект репозитория чатов
func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) CreateChat(ctx context.Context, in *model.CreateChatRequest) (int64, error) {
	builder := sq.Insert(chatTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn).
		Values(in.Title).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create_InsertChatTable",
		QueryRaw: query,
	}

	var chatID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	builder = sq.Insert(chatUsersTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, userTagColumn)

	for _, v := range in.UserTags {
		builder = builder.Values(chatID, v)
	}

	query, args, err = builder.ToSql()
	if err != nil {
		return 0, err
	}

	q = db.Query{
		Name:     "chat_repository.Create_InsertChatUserTable",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

func (r *repo) AddUserToChat(ctx context.Context, in *model.AddUserToChatRequest) error {
	builder := sq.Insert(chatUsersTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, userTagColumn)
	for _, v := range in.UserTags {
		builder = builder.Values(in.ChatID, v)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.AddUserToChat_Insert",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DeleteChat(ctx context.Context, chatID int64) error {
	builder := sq.Delete(messagesTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIDColumn: chatID})
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete_DeleteMessagesFromChat",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	builder = sq.Delete(chatUsersTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIDColumn: chatID})
	query, args, err = builder.ToSql()
	if err != nil {
		return err
	}

	q = db.Query{
		Name:     "chat_repository.Delete_DeleteUsersFromChat",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	builder = sq.Delete(chatTableName).
		Where(sq.Eq{idColumn: chatID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err = builder.ToSql()
	if err != nil {
		return err
	}

	q = db.Query{
		Name:     "chat_repository.Delete_DeleteChat",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) SendMessageToChat(ctx context.Context, in *model.SendMessageToChatRequest) error {
	builder := sq.Insert(messagesTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, userTagColumn, messageColumn).
		Values(in.ChatID, in.UserTag, in.Message)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.SendMessageToChat_Insert",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}
