package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"chat_server/internal/converter"
	desc "chat_server/pkg/chat_server_v1"
)

// AddUserToChat является апи методом для добавления пользователя в чат
func (i *Implementation) AddUserToChat(ctx context.Context, in *desc.AddUsersToChatRequest) (*emptypb.Empty, error) {
	err := i.chatService.AddUser(ctx, converter.ToAddUserToChatRequestFromAPI(in))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
