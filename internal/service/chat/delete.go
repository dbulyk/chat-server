package chat

import (
	"context"
)

// Delete является сервисной прослойкой для удаления чата
func (s *service) Delete(ctx context.Context, chatID int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		txErr = s.chatRepository.DeleteChat(ctx, chatID)
		return txErr
	})
	return err
}
