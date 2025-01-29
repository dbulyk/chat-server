package chat

import (
	"context"
)

// Delete является сервисной прослойкой для удаления чата
func (s *service) Delete(ctx context.Context, chatID int64) error {
	err := s.chatRepository.DeleteChat(ctx, chatID)
	if err != nil {
		return err
	}
	return nil
}
