package service

import (
	"context"

	"chat_server/internal/model"
	desc "chat_server/pkg/chat_server_v1"
)

// ChatService подключает все функции сервиса в один интерфейс
type ChatService interface {
	CreateChatServ(ctx context.Context, chatInfo *model.CreateChat) (int64, error)
	DeleteChatServ(ctx context.Context, chatID int64) error
	AddMembersServ(ctx context.Context, chatID int64, memberTags []string) error
	RemoveMembersServ(ctx context.Context, chatID int64, memberTags []string) error
	SendMessageServ(ctx context.Context, message *desc.Message) error
}
