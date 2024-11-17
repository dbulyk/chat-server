package chat

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"chat_server/internal/repository"
)

var _ repository.ChatServerRepository = (*Repo)(nil)

// Repo объединяет модули приложения
type Repo struct {
	repoChat
	repoMember
	repoMessage
}

// NewRepository возвращает репозиторий с имплементациями основных функций приложения
func NewRepository(db *pgxpool.Pool) *Repo {
	return &Repo{
		repoChat:    repoChat{db: db},
		repoMember:  repoMember{db: db},
		repoMessage: repoMessage{db: db},
	}
}
