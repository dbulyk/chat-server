package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"

	"chat_server/internal/repository"
	"chat_server/internal/repository/chat/model"
)

var _ repository.Message = (*repoMessage)(nil)

type repoMessage struct {
	db *pgxpool.Pool
}

// SendMessage отправляет сообщение в чат
func (r *repoMessage) SendMessage(ctx context.Context, msg model.Message) error {
	builder := sq.Insert("messages").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_tag", "message").
		Values(msg.ChatID, msg.MemberTag, msg.Text)

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
