package chat

import (
	"chat_server/internal/client/db"
	"chat_server/internal/repository"
	serv "chat_server/internal/service"
)

type service struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

// NewChatService возвращает объект сервиса чатов
func NewChatService(chatRepository repository.ChatRepository, txManager db.TxManager) serv.ChatService {
	return &service{chatRepository: chatRepository, txManager: txManager}
}
