package chat

import (
	"github.com/dbulyk/platform_common/pkg/db"

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
