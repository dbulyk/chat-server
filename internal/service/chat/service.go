package chat

import (
	"chat_server/internal/repository"
	serv "chat_server/internal/service"
)

type service struct {
	chatRepository repository.ChatRepository
}

// NewChatService возвращает объект сервиса чатов
func NewChatService(chatRepository repository.ChatRepository) serv.ChatService {
	return &service{chatRepository: chatRepository}
}
