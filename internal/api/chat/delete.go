package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "chat_server/pkg/chat_server_v1"
)

// DeleteChat является апи методом для удаления чата
func (i *Implementation) DeleteChat(ctx context.Context, in *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	err := i.chatService.Delete(ctx, in.GetChatId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
