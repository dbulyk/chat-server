package chat

import (
	"chat_server/internal/model"
	"context"
)

func (s *Service) CreateChat(ctx context.Context, in *model.CreateChatRequest) (*model.CreateChatResponse, error) {
	chatID, err := s.chatRepository.CreateChat(ctx, in)
	if err != nil {
		return nil, err
	}
	return &model.CreateChatResponse{ChatId: chatID}, nil
}
