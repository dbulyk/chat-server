package repository

import (
	"chat_server/internal/repository/chat/model"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatRepository interface {
	CreateChat(ctx context.Context, in *model.CreateChatRequest) (*model.CreateChatResponse, error)
	AddUserToChat(ctx context.Context, in *model.AddUserToChatRequest) (*emptypb.Empty, error)
	DeleteChat(ctx context.Context, in *model.DeleteChatRequest) (*emptypb.Empty, error)
	SendMessageToChat(ctx context.Context, in *model.SendMessageToChatRequest) (*emptypb.Empty, error)
}
