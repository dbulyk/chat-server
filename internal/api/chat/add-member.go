package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "chat_server/pkg/chat_server_v1"
)

// AddMembersToChat добавляет пользователей в уже созданный чат
func (i *Implementation) AddMembersToChat(ctx context.Context, in *desc.AddUsersToChatRequest) (*emptypb.Empty, error) {
	err := i.chatService.AddMembersServ(ctx, in.GetChatId(), in.GetUsersTag())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
