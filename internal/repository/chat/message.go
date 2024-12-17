package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"chat_server/internal/client/db"
	"chat_server/internal/model"
	"chat_server/internal/repository"
)

var _ repository.Message = (*repoMessage)(nil)

type repoMessage struct {
	db db.Client
}

// SendMessage отправляет сообщение в чат
func (r *repoMessage) SendMessage(ctx context.Context, msg *model.Message) error {
	builder := sq.Insert("messages").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_tag", "message").
		Values(msg.ChatID, msg.MemberTag, msg.Text)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.SendMessage",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}
