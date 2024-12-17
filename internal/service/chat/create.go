package chat

import (
	"context"

	"chat_server/internal/model"
)

// CreateChatServ является сервисной прослойкой для создания чата
func (s *serv) CreateChatServ(ctx context.Context, chatInfo *model.CreateChat) (int64, error) {
	res, err := s.chatServerRepository.CreateChat(ctx, chatInfo)
	if err != nil {
		return 0, err
	}
	return res, nil
}
