package chat

import (
	"chat_server/internal/repository"
	"chat_server/internal/service"
)

var _ service.ChatService = (*serv)(nil)

type serv struct {
	chatServerRepository repository.ChatServerRepository
}

// NewChatService создаёт новый объект сервиса чатов
func NewChatService(chatServerRepository repository.ChatServerRepository) *serv {
	return &serv{chatServerRepository: chatServerRepository}
}
