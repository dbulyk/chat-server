package chat

import (
	"context"

	"chat_server/internal/model"
)

// AddUser является сервисной прослойкой для добавления пользователя в чат
func (s *service) AddUser(ctx context.Context, in *model.AddUserToChatRequest) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.chatRepository.AddUserToChat(ctx, in)
		return errTx
	})
	return err
}
