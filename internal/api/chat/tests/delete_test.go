package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"chat_server/internal/api/chat"
	"chat_server/internal/service"
	"chat_server/internal/service/mocks"
	desc "chat_server/pkg/chat_server_v1"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.DeleteChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID = gofakeit.Int64()

		req = &desc.DeleteChatRequest{ChatId: chatID}

		serviceErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "Успешное удаление чата",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.DeleteMock.Expect(ctx, chatID).Return(nil)
				return mock
			},
		},
		{
			name: "Ошибка удаления чата",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.DeleteMock.Expect(ctx, chatID).Return(serviceErr)
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewImplementation(chatServiceMock)

			newID, err := api.DeleteChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
