package chat

import (
	"chat_server/internal/client/db"
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
func NewRepository(db db.Client) *Repo {
	return &Repo{
		repoChat:    repoChat{db: db},
		repoMember:  repoMember{db: db},
		repoMessage: repoMessage{db: db},
	}
}
