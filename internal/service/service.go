package service

import (
	"context"

	"chat_server/internal/model"
)

// ChatService описывает контракт для сервиса чатов
type ChatService interface {
	Create(ctx context.Context, in *model.CreateChatRequest) (int64, error)
	AddUser(ctx context.Context, in *model.AddUserToChatRequest) error
	Delete(ctx context.Context, chatID int64) error
	SendMessage(ctx context.Context, in *model.SendMessageToChatRequest) error
}
