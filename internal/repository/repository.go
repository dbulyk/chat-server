package repository

import (
	"context"

	"chat_server/internal/model"
)

// ChatRepository описывает контракт репозитория
type ChatRepository interface {
	CreateChat(ctx context.Context, in *model.CreateChatRequest) (int64, error)
	AddUserToChat(ctx context.Context, in *model.AddUserToChatRequest) error
	DeleteChat(ctx context.Context, chatID int64) error
	SendMessageToChat(ctx context.Context, in *model.SendMessageToChatRequest) error
}
