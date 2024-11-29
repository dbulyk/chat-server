package converter

import (
	"chat_server/internal/model"
	desc "chat_server/pkg/chat_server_v1"
)

// ToMessageFromDesc конвертирует модель из прото в сервисную модель
func ToMessageFromDesc(req *desc.Message) model.Message {
	return model.Message{
		ChatID:    req.GetChatId(),
		Text:      req.GetText(),
		MemberTag: req.GetUserTag(),
	}
}
