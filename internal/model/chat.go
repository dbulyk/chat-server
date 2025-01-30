package model

// CreateChatRequest является структурой для запроса создания чата
type CreateChatRequest struct {
	Title    string
	UserTags []string
}

// CreateChatResponse является структурой для результатов создания чата
type CreateChatResponse struct {
	ChatID int64
}

// DeleteChatRequest является структурой для запроса удаления чата
type DeleteChatRequest struct {
	ChatID int64
}

// AddUserToChatRequest является структурой для запроса добавления пользователя в чат
type AddUserToChatRequest struct {
	ChatID   int64
	UserTags []string
}

// SendMessageToChatRequest является структурой для запроса отправки сообщения в чат
type SendMessageToChatRequest struct {
	ChatID  int64
	Message string
	UserTag string
}
