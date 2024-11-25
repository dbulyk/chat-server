package chat

import (
	"context"

	"chat_server/internal/converter"
	desc "chat_server/pkg/chat_server_v1"
)

// SendMessageServ является сервисной прослойкой для отправки сообщения в чат
func (s *serv) SendMessageServ(ctx context.Context, msg *desc.Message) error {
	message := converter.ToMessageFromDesc(msg)
	err := s.chatServerRepository.SendMessage(ctx, &message)
	if err != nil {
		return err
	}
	return nil
}
