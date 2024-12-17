package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"chat_server/internal/client/db"
	"chat_server/internal/repository"
)

var _ repository.Member = (*repoMember)(nil)

type repoMember struct {
	db db.Client
}

// AddMembers добавляет пользователя(ей) в чат
func (r *repoMember) AddMembers(ctx context.Context, chatID int64, memberTags []string) error {
	builder := sq.Insert("users_chats").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_tag")
	for _, v := range memberTags {
		builder = builder.Values(chatID, v)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.AddMembers",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}

// RemoveMembers удаляет пользователя(ей) из чата
func (r *repoMember) RemoveMembers(ctx context.Context, chatID int64, memberTags []string) error {
	builder := sq.Delete("users_chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.And{
			sq.Eq{"chat_id": chatID},
			sq.Eq{"user_tag": memberTags},
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.AddMembers",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
