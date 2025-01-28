package chat

import "chat_server/internal/repository"

type Service struct {
	chatRepository repository.ChatRepository
}

func NewChatService(chatRepository repository.ChatRepository) *Service {
	return &Service{chatRepository: chatRepository}
}
