package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"chat_server/internal/api/chat"
	"chat_server/internal/model"
	"chat_server/internal/service"
	"chat_server/internal/service/mocks"
	desc "chat_server/pkg/chat_server_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.CreateChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		title    = gofakeit.BookTitle()
		userTags = []string{
			gofakeit.Gamertag(),
			gofakeit.Gamertag(),
		}

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateChatRequest{
			Title:     title,
			UsersTags: userTags,
		}

		modelReq = &model.CreateChatRequest{
			Title:    title,
			UserTags: userTags,
		}

		res = &desc.CreateChatResponse{
			ChatId: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateChatResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "Успешное создание чата",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, modelReq).Return(id, nil)
				return mock
			},
		},
		{
			name: "Ошибка создания чата",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, modelReq).Return(0, serviceErr)
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

			newID, err := api.CreateChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
