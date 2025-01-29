package chat

import (
	"context"

	"chat_server/internal/model"
)

// Create является сервисной прослойкой для создания чата
func (s *service) Create(ctx context.Context, in *model.CreateChatRequest) (int64, error) {
	chatID, err := s.chatRepository.CreateChat(ctx, in)
	if err != nil {
		return 0, err
	}
	return chatID, nil
}
