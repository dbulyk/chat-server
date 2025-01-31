package chat

import (
	"context"

	"chat_server/internal/model"
)

// SendMessage является сервисной прослойкой для отправки сообщения в чат
func (s *service) SendMessage(ctx context.Context, in *model.SendMessageToChatRequest) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		txErr = s.chatRepository.SendMessageToChat(ctx, in)
		return txErr
	})
	return err
}
