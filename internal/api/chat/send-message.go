package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "chat_server/pkg/chat_server_v1"
)

// SendMessage отправляет сообщение в чат
func (i *Implementation) SendMessage(ctx context.Context, in *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.chatService.SendMessageServ(ctx, in.GetMessage())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
