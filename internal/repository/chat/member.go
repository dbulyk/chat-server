package chat

import (
	"chat_server/internal/repository"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ repository.Member = (*repoMember)(nil)

type repoMember struct {
	db *pgxpool.Pool
}

// AddMembersToChat добавляет пользователя(ей) в чат
func (r *repoMember) AddMembersToChat(ctx context.Context, chatID int64, memberTags []string) error {
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
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

// RemoveMembersFromChat удаляет пользователя(ей) из чата
func (r *repoMember) RemoveMembersFromChat(ctx context.Context, chatID int64, memberTags []string) error {
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

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
