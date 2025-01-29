package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"chat_server/internal/converter"
	desc "chat_server/pkg/chat_server_v1"
)

// SendMessage является апи методом для отправки сообщения в чат
func (i *Implementation) SendMessage(ctx context.Context, in *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.chatService.SendMessage(ctx, converter.SendMessageRequestFromAPI(in))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
