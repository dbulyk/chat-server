package converter

import (
	"chat_server/internal/model"
	desc "chat_server/pkg/chat_server_v1"
)

// ToCreateChatRequestFromAPI конвертирует модель создания чата из протобафа в сервисную
func ToCreateChatRequestFromAPI(in *desc.CreateChatRequest) *model.CreateChatRequest {
	return &model.CreateChatRequest{
		Title:    in.GetTitle(),
		UserTags: in.GetUsersTags(),
	}
}

// ToAddUserToChatRequestFromAPI конвертирует модель добавления пользователя в чат из протобафа в сервисную
func ToAddUserToChatRequestFromAPI(in *desc.AddUsersToChatRequest) *model.AddUserToChatRequest {
	return &model.AddUserToChatRequest{
		ChatID:   in.GetChatId(),
		UserTags: in.GetUsersTag(),
	}
}

// SendMessageRequestFromAPI конвертирует модель отправки сообщения из протобафа в сервисную
func SendMessageRequestFromAPI(in *desc.SendMessageRequest) *model.SendMessageToChatRequest {
	return &model.SendMessageToChatRequest{
		ChatID:  in.GetChatId(),
		Message: in.GetText(),
		UserTag: in.GetUserTag(),
	}
}
