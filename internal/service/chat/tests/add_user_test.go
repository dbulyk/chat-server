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
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestAddUser(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.AddUserToChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID   = gofakeit.Int64()
		userTags = []string{gofakeit.Word(), gofakeit.Word()}

		req     = &model.AddUserToChatRequest{ChatID: chatID, UserTags: userTags}
		repoErr = fmt.Errorf("repo error")
	)

	tests := []struct {
		name          string
		args          args
		want          *emptypb.Empty
		err           error
		chatRepoMock  chatRepositoryMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "Успешное добавление пользователя в чат",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := mocks.NewChatRepositoryMock(mc)
				mock.AddUserToChatMock.Expect(ctx, req).Return(nil)
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
			name: "Ошибка добавления пользователя в чат",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: repoErr,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := mocks.NewChatRepositoryMock(mc)
				mock.AddUserToChatMock.Expect(ctx, req).Return(repoErr)
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

			err := serv.AddUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
