package repository

import (
	"context"

	"chat_server/internal/model"
)

// ChatServerRepository определяет взаимодействие с бд
type ChatServerRepository interface {
	Chat
	Member
	Message
}

// Chat определяет взаимодействие с чатом
type Chat interface {
	CreateChat(ctx context.Context, chatInfo *model.CreateChat) (int64, error)
	DeleteChat(ctx context.Context, chatID int64) error
}

// Member определяет взаимодействие с участниками чата
type Member interface {
	AddMembers(ctx context.Context, chatID int64, memberTags []string) error
	RemoveMembers(ctx context.Context, chatID int64, memberTags []string) error
}

// Message определяет взаимодействие с сообщениями
type Message interface {
	SendMessage(ctx context.Context, msg *model.Message) error
}
