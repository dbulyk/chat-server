package chat

import (
	"chat_server/internal/service"
	desc "chat_server/pkg/chat_server_v1"
)

// Implementation представляет реализацию сервиса для работы с чатами
type Implementation struct {
	desc.UnimplementedChatServerV1Server
	chatService service.ChatService
}

// NewImplementation возвращает объект сервиса для работы с чатом
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{chatService: chatService}
}
