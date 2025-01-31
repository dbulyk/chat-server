package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dbulyk/platform_common/pkg/db"
	mocks2 "github.com/dbulyk/platform_common/pkg/db/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"chat_server/internal/model"
	"chat_server/internal/repository"
	"chat_server/internal/repository/mocks"
	"chat_server/internal/service/chat"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.SendMessageToChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		text    = gofakeit.Letter()
		userTag = gofakeit.Gamertag()
		chatID  = gofakeit.Int64()

		req     = &model.SendMessageToChatRequest{ChatID: chatID, UserTag: userTag, Message: text}
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
				mock.SendMessageToChatMock.Expect(ctx, req).Return(nil)
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
				mock.SendMessageToChatMock.Expect(ctx, req).Return(repoErr)
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

			err := serv.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
