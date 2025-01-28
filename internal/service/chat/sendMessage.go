package chat

import (
	"chat_server/internal/model"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) SendMessageToChat(ctx context.Context, in *model.SendMessageToChatRequest) (*emptypb.Empty, error) {
	_, err := s.chatRepository.SendMessageToChat(ctx, in)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
