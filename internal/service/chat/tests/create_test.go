package tests

import (
	"chat_server/internal/client/db"
	mocks2 "chat_server/internal/client/db/mocks"
	"chat_server/internal/model"
	"chat_server/internal/repository"
	"chat_server/internal/repository/mocks"
	"chat_server/internal/service/chat"
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.CreateChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		title    = gofakeit.Name()
		userTags = []string{gofakeit.Gamertag(), gofakeit.Gamertag()}
		chatID   = gofakeit.Int64()

		req     = &model.CreateChatRequest{Title: title, UserTags: userTags}
		repoErr = fmt.Errorf("repo error")
	)

	tests := []struct {
		name          string
		args          args
		want          int64
		err           error
		chatRepoMock  chatRepositoryMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "Успешное создание чата",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: chatID,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := mocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(chatID, nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := mocks2.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
		},
		{
			name: "Ошибка создания чата",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: repoErr,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := mocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(0, repoErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := mocks2.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepoMock := tt.chatRepoMock(mc)
			txManagerMock := tt.txManagerMock(mc)
			serv := chat.NewChatService(chatRepoMock, txManagerMock)

			chatId, err := serv.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, chatId)
		})
	}
}
