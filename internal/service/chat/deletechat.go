package chat

import (
	"context"
)

// DeleteChatServ является сервисной прослойкой для удаления чата
func (s *serv) DeleteChatServ(ctx context.Context, chatID int64) error {
	err := s.chatServerRepository.DeleteChat(ctx, chatID)
	if err != nil {
		return err
	}
	return nil
}
