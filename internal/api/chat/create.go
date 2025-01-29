package chat

import (
	"context"

	"chat_server/internal/converter"
	desc "chat_server/pkg/chat_server_v1"
)

// CreateChat является апи методом для создания чата
func (i *Implementation) CreateChat(ctx context.Context, in *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	chatID, err := i.chatService.Create(ctx, converter.ToCreateChatRequestFromAPI(in))
	if err != nil {
		return nil, err
	}
	return &desc.CreateChatResponse{ChatId: chatID}, nil
}
