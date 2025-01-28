package repository

import (
	"chat_server/internal/model"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatRepository interface {
	CreateChat(ctx context.Context, in *model.CreateChatRequest) (int64, error)
	AddUserToChat(ctx context.Context, in *model.AddUserToChatRequest) (*emptypb.Empty, error)
	DeleteChat(ctx context.Context, chatId int64) (*emptypb.Empty, error)
	SendMessageToChat(ctx context.Context, in *model.SendMessageToChatRequest) (*emptypb.Empty, error)
}
