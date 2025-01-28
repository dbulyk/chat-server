package chat

import (
	"chat_server/internal/model"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) DeleteChat(ctx context.Context, in *model.DeleteChatRequest) (*emptypb.Empty, error) {
	_, err := s.chatRepository.DeleteChat(ctx, in.ChatId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
