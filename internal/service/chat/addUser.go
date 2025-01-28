package chat

import (
	"chat_server/internal/model"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) AddUserToChat(ctx context.Context, in *model.AddUserToChatRequest) (*emptypb.Empty, error) {
	_, err := s.chatRepository.AddUserToChat(ctx, in)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
