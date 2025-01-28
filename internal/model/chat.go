package model

type CreateChatRequest struct {
	Title    string
	UserTags []string
}

type CreateChatResponse struct {
	ChatId int64
}

type DeleteChatRequest struct {
	ChatId int64
}

type AddUserToChatRequest struct {
	ChatId   int64
	UserTags []string
}

type SendMessageToChatRequest struct {
	ChatId  int64
	Message string
	UserTag string
}
