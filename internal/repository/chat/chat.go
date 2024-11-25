package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"

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
	db *pgxpool.Pool
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

	var chatID int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&chatID)
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

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
