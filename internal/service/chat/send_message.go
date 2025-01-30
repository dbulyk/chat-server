package chat

import (
	"context"

	"chat_server/internal/model"
)

// SendMessage является сервисной прослойкой для отправки сообщения в чат
func (s *service) SendMessage(ctx context.Context, in *model.SendMessageToChatRequest) error {
	err := s.chatRepository.SendMessageToChat(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
