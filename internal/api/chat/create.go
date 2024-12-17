package chat

import (
	"context"

	"chat_server/internal/model"
	desc "chat_server/pkg/chat_server_v1"
)

// CreateChat Создаёт новый чат с указанным названием
func (i *Implementation) CreateChat(ctx context.Context, in *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	createChatModel := model.CreateChat{Title: in.GetTitle()}
	chatID, err := i.chatService.CreateChatServ(ctx, &createChatModel)
	if err != nil {
		return nil, err
	}

	return &desc.CreateChatResponse{ChatId: chatID}, nil
}
