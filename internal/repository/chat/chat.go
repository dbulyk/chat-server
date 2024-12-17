package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"chat_server/internal/client/db"
	"chat_server/internal/model"
	"chat_server/internal/repository"
)

var _ repository.Chat = (*repoChat)(nil)

const (
	tableName = "chats"

	idColumn        = "id"
	titleColumn     = "title"
	createdAtColumn = "created_at"
)

type repoChat struct {
	db db.Client
}

// CreateChat создаёт чат с заданным названием
func (r *repoChat) CreateChat(ctx context.Context, chatInfo *model.CreateChat) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn).
		Values(chatInfo.Title).
		Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return -1, err
	}

	q := db.Query{
		Name:     "chat_repository.CreateChat",
		QueryRaw: query,
	}

	var chatID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return -1, err
	}

	return chatID, nil
}

// DeleteChat удаляет чат
func (r *repoChat) DeleteChat(ctx context.Context, chatID int64) error {
	builder := sq.Delete(tableName).
		Where(sq.Eq{idColumn: chatID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.DeleteChat",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}
