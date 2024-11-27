package chat

import (
	"chat_server/internal/service"
	desc "chat_server/pkg/chat_server_v1"
)

// Implementation является объектом сервера
type Implementation struct {
	desc.UnimplementedChatServerV1Server
	chatService service.ChatService
}

// NewImplementation создаёт объект сервера
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
