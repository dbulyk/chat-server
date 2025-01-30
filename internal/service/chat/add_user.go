package chat

import (
	"context"

	"chat_server/internal/model"
)

// AddUser является сервисной прослойкой для добавления пользователя в чат
func (s *service) AddUser(ctx context.Context, in *model.AddUserToChatRequest) error {
	err := s.chatRepository.AddUserToChat(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
