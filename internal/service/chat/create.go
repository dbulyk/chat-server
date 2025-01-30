package chat

import (
	"context"

	"chat_server/internal/model"
)

// Create является сервисной прослойкой для создания чата
func (s *service) Create(ctx context.Context, in *model.CreateChatRequest) (int64, error) {
	var chatID int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		chatID, errTx = s.chatRepository.CreateChat(ctx, in)
		if errTx != nil {
			return errTx
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return chatID, nil
}
